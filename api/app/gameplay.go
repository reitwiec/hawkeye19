package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type CheckAnswer struct {
	Answer   string
	RegionId int
	Level    int
}

type Stats struct {
	TotalPlayers   int
	AnswerAttempts int
	MaxPoints      int
	SameLevel      int
	Leading        int
	Trailing       int
}

const (
	RegionComplete = 4 //no of questions + 1
	LinearRegionId = 6
)

func (hawk *App) checkAnswer(w http.ResponseWriter, r *http.Request) {
	//obtain currUser from context
	currUser := r.Context().Value("User").(User)
	//obtain answer, regionId from request
	checkAns := CheckAnswer{}
	err := json.NewDecoder(r.Body).Decode(&checkAns)
	if err != nil {
		fmt.Println("Could not decode checkAnswer struct " + err.Error())
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusBadRequest, w)
		return
	}
	//get level from currUser
	switch checkAns.RegionId {
	case 1:
		checkAns.Level = currUser.Region1
	case 2:
		checkAns.Level = currUser.Region2
	case 3:
		checkAns.Level = currUser.Region3
	case 4:
		checkAns.Level = currUser.Region4
	case 5:
		checkAns.Level = currUser.Region5
	case 6:
		checkAns.Level = currUser.Region6

	}

	//sanitize answer
	checkAns.Answer = Sanitize(checkAns.Answer)

	//find actual answer
	question := Question{}
	err = hawk.DB.Where("level = ? AND region = ?", checkAns.Level, checkAns.RegionId).First(&question).Error
	actualAns := question.Answer
	//check if same, close or wrong
	status := CheckAnswerStatus(checkAns.Answer, actualAns)
	//log attempt
	attempt := Attempt{
		User:      currUser.ID,
		Question:  question.ID,
		Answer:    checkAns.Answer,
		Status:    status,
		Timestamp: time.Now(),
	}
	tx := hawk.DB.Begin()
	err = tx.Create(&attempt).Error
	if err != nil {
		fmt.Println("Could not log answer attempt")
		ResponseWriter(false, "Could not log answer", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	//if answer is correct
	if status == CorrectAnswer {
		//answer is correct
		//update currUser level
		tx := hawk.DB.Begin()
		region := "Region" + strconv.Itoa(checkAns.RegionId)
		err = tx.Model(&currUser).Update(region, checkAns.Level+1).Error
		if err != nil {
			fmt.Println("Could not update level")
			ResponseWriter(false, "Could not update level", nil, http.StatusInternalServerError, w)
			tx.Rollback()
		}

		//unlock region
		isRegionComplete := false
		switch checkAns.RegionId {
		case 1:
			if currUser.Region1 == RegionComplete {
				isRegionComplete = true
			}
		case 2:
			if currUser.Region2 == RegionComplete {
				isRegionComplete = true

			}
		case 3:
			if currUser.Region3 == RegionComplete {
				isRegionComplete = true
			}
		case 4:
			if currUser.Region4 == RegionComplete {
				isRegionComplete = true
			}
		case 5:
			if currUser.Region5 == RegionComplete {
				//unlock linear gameplay
				err = tx.Model(&currUser).Update("Linear", 1).Error
				if err != nil {
					fmt.Println("Could not unlock linear region")
					ResponseWriter(false, "Could not unlock linear region", nil, http.StatusInternalServerError, w)
					tx.Rollback()
					return
				}

			}
		}
		if isRegionComplete {

			//get next region
			nextRegion := GetNextRegion(currUser.UnlockOrder)
			currUser.UnlockOrder = UpdateUnlockOrder(currUser.UnlockOrder, int(nextRegion)-48)
			//update DB with unlocked region
			switch nextRegion {
			case '2':
				err = tx.Model(&currUser).Update("Region2", 1).Error
			case '3':
				err = tx.Model(&currUser).Update("Region3", 1).Error
			case '4':
				err = tx.Model(&currUser).Update("Region4", 1).Error
			case '5':
				err = tx.Model(&currUser).Update("Region5", 1).Error
			}
			if err != nil {
				fmt.Println("Could not unlock region in DB")
				ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
				tx.Rollback()
				return
			}
			//update UnlockOrder string in DB
			err = tx.Model(&currUser).Update("unlock_order", currUser.UnlockOrder).Error
			if err != nil {
				fmt.Println("Could not unlock region in DB")
				ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
				tx.Rollback()
				return
			}
		}
		tx.Commit()

	}
	ResponseWriter(true, "Answer status", status, http.StatusOK, w)
}

func (hawk *App) getRecentTries(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)
	//get user ID from context, region and level from GET request
	keys, ok := r.URL.Query()["question"]
	if !ok || len(keys) < 1 {
		fmt.Println("Param absent")
		ResponseWriter(false, "Param absent", nil, http.StatusBadRequest, w)
		return
	}
	var answers []string
	//query db for answers
	err := hawk.DB.Model(&Attempt{}).Where("question = ? AND user = ?", keys[0], currUser.ID).Order("timestamp desc").Pluck("answer", &answers).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		fmt.Println("Database error")
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "", answers, http.StatusOK, w)
}

func (hawk *App) getQuestion(w http.ResponseWriter, r *http.Request) {

	currUser := r.Context().Value("User").(User)
	//@TODO: Implement in middleware
	if (currUser == User{}) {
		ResponseWriter(false, "User not logged in.", nil, http.StatusNetworkAuthenticationRequired, w)
		return
	}

	keys, ok := r.URL.Query()["region"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Invalid request", nil, http.StatusBadRequest, w)
		return
	}

	key := keys[0]
	level := 0

	switch key {
	case "1":
		level = currUser.Region1
	case "2":
		level = currUser.Region2
	case "3":
		level = currUser.Region3
	case "4":
		level = currUser.Region4
	case "5":
		level = currUser.Region5
	}

	question := Question{}

	err := hawk.DB.Select("id, question, level, region, add_info").Where("level = ? AND region = ?", level, key).First(&question).Error
	if err != nil {
		ResponseWriter(false, "Could not fetch question.", nil, http.StatusInternalServerError, w)
		return
	}

	ResponseWriter(true, "Question fetched.", question, http.StatusOK, w)
}

func (hawk *App) getHints(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["question"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Invalid request", nil, http.StatusBadRequest, w)
		return
	}

	key := keys[0]

	var hints []string

	err := hawk.DB.Model(&Hint{}).Where("question=? AND active=1", key).Pluck("hint", &hints).Error
	if err != nil {
		ResponseWriter(false, "Could not fetch hint.", nil, http.StatusInternalServerError, w)
		return
	}

	ResponseWriter(true, "Hints fetched.", hints, http.StatusOK, w)
}

func (hawk *App) getStats(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)
	if (currUser == User{}) {
		ResponseWriter(false, "User not logged in.", nil, http.StatusNetworkAuthenticationRequired, w)
		return
	}

	currStats := Stats{}

	err := hawk.DB.Model(&User{}).Count(&currStats.TotalPlayers).Error
	if err != nil {
		ResponseWriter(false, "Cant get total players", nil, http.StatusInternalServerError, w)
		return
	}

	err = hawk.DB.Model(&User{}).Where("points = ?", currUser.Points).Count(&currStats.SameLevel).Error
	if err != nil {
		ResponseWriter(false, "Cant get players at par", nil, http.StatusInternalServerError, w)
		return
	}

	err = hawk.DB.Model(&User{}).Where("points > ?", currUser.Points).Count(&currStats.Leading).Error
	if err != nil {
		ResponseWriter(false, "Cant get leading users", nil, http.StatusInternalServerError, w)
		return
	}

	currStats.TotalPlayers += 1
	currStats.Leading += 1
	currStats.Trailing = currStats.TotalPlayers - (currStats.Leading + currStats.SameLevel)

	ResponseWriter(true, "Current stats", currStats, http.StatusOK, w)

}

func (hawk *App) getSideQuestQuestion(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)
	//get sidequest level from request
	keys, ok := r.URL.Query()["level"]
	if !ok || len(keys) < 1 {
		fmt.Println("Params missing")
		ResponseWriter(false, "Param missing", nil, http.StatusBadRequest, w)
		return
	}
	index, err := strconv.Atoi(keys[0])
	index -= 1
	if err != nil {
		fmt.Println("Error in Atoi")
		ResponseWriter(false, "Error in Atoi", nil, http.StatusInternalServerError, w)
		return
	}
	level := currUser.SideQuest[index] - 48 //ascii to int
	fmt.Println(level)
	question := Question{}
	err = hawk.DB.Where("region = ? AND level = ?", 0, level).First(&question).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Question does not exist")
			ResponseWriter(false, "Question does not exist", nil, http.StatusNotFound, w)
			return
		} else {
			fmt.Println("Database error")
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}

	}
	question.Answer = ""
	ResponseWriter(true, "Question fetched", question, http.StatusOK, w)
}
