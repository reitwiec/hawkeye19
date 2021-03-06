package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type CheckAnswer struct {
	Answer   string `json:"answer"`
	RegionID int    `json:"regionID" validate:"min=0,max=5"`
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
	RegionComplete = 5 //no of questions + 1
)

func (hawk *App) checkAnswer(w http.ResponseWriter, r *http.Request) {
	//obtain currUser from context
	currUser := r.Context().Value("User").(User)
	// check if verified and not banned
	if currUser.IsVerified != 1 {
		ResponseWriter(false, "Not verified", nil, http.StatusOK, w)
		return
	}
	if currUser.Banned == 1 {
		ResponseWriter(false, "User banned", nil, http.StatusOK, w)
		return
	}
	//obtain answer, regionId from request
	checkAns := CheckAnswer{}
	err := json.NewDecoder(r.Body).Decode(&checkAns)
	if err != nil {
		fmt.Println("Could not decode checkAnswer struct " + err.Error())
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusBadRequest, w)
		return
	}

	level := -1

	//get level from currUser
	switch checkAns.RegionID {
	case 1:
		level = currUser.Region1
	case 2:
		level = currUser.Region2
	case 3:
		level = currUser.Region3
	case 4:
		level = currUser.Region4
	case 5:
		level = currUser.Region5

	}

	if level == RegionComplete {
		ResponseWriter(false, "Region has been completed", nil, http.StatusBadRequest, w)
		return
	}

	//sanitize answer
	checkAns.Answer = SanitizeAnswer(checkAns.Answer)
	//find actual answer
	question := Question{}
	err = hawk.DB.Where("level = ? AND region = ?", level, checkAns.RegionID).First(&question).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Question not found", nil, http.StatusNotFound, w)
			return
		}
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
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
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not log answer", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}

	//if answer is correct
	if status == CorrectAnswer {
		//answer is correct
		//update currUser level

		switch checkAns.RegionID {
		case 1:
			err = tx.Model(&currUser).Updates(User{Region1: level + 1, Points: currUser.Points + 1}).Error

		case 2:
			err = tx.Model(&currUser).Updates(User{Region2: level + 1, Points: currUser.Points + 1}).Error

		case 3:
			err = tx.Model(&currUser).Updates(User{Region3: level + 1, Points: currUser.Points + 1}).Error

		case 4:
			err = tx.Model(&currUser).Updates(User{Region4: level + 1, Points: currUser.Points + 1}).Error

		case 5:
			err = tx.Model(&currUser).Updates(User{Region5: level + 1, Points: currUser.Points + 1}).Error
		}

		if err != nil {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Could not update level", nil, http.StatusInternalServerError, w)
			tx.Rollback()
		}

		//unlock region
		isRegionComplete := false
		switch checkAns.RegionID {
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
			/*
				case 5:
					if currUser.Region5 == RegionComplete {
						//unlock linear gameplay
						err = tx.Model(&currUser).Update("Region6", 1).Error
						if err != nil {
							fmt.Println("Could not unlock linear region")
							ResponseWriter(false, "Could not unlock linear region", nil, http.StatusInternalServerError, w)
							tx.Rollback()
							return
						}

					}
			*/
		}
		if isRegionComplete {
			//is it time to unlock linear ?
			if currUser.Region5 == 0 && currUser.Region1+currUser.Region2+currUser.Region3+currUser.Region4 == RegionComplete*4 {
				err = tx.Model(&currUser).Update("Region5", 1).Error
				if err != nil {
					LogRequest(r, ERROR, err.Error())
					ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
					tx.Rollback()
					return
				}
				//update UnlockOrder string in DB
				err = tx.Model(&currUser).Update("unlock_order", "0,0,0,0,0").Error
			} else {
				//get next region
				nextRegion := int(GetNextRegion(currUser.UnlockOrder)) - 48
				currUser.UnlockOrder = UpdateUnlockOrder(currUser.UnlockOrder, nextRegion)
				//update DB with unlocked region
				switch nextRegion {
				case 2:
					err = tx.Model(&currUser).Update("Region2", 1).Error
				case 3:
					err = tx.Model(&currUser).Update("Region3", 1).Error
				case 4:
					err = tx.Model(&currUser).Update("Region4", 1).Error
					//case 5: //linear
					//err = tx.Model(&currUser).Update("Region5", 1).Error
				}

				if err != nil {
					LogRequest(r, ERROR, err.Error())
					ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
					tx.Rollback()
					return
				}

				//update UnlockOrder string in DB
				err = tx.Model(&currUser).Update("unlock_order", currUser.UnlockOrder).Error
				if err != nil {
					LogRequest(r, ERROR, err.Error())
					ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
					tx.Rollback()
					return
				}
			}
		}

	}
	tx.Commit()
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
	err := hawk.DB.Model(&Attempt{}).Where("question = ? AND user = ?", keys[0], currUser.ID).Order("timestamp desc").Limit(7).Pluck("answer", &answers).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		LogRequest(r, ERROR, err.Error())
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
		if gorm.IsRecordNotFoundError(err) {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Question doesn't exist", nil, http.StatusInternalServerError, w)
			fmt.Println(err)
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Could not fetch question.", nil, http.StatusInternalServerError, w)
			fmt.Println(err)
		}
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
		if gorm.IsRecordNotFoundError(err) {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Question doesn't exist", nil, http.StatusInternalServerError, w)
			fmt.Println(err)
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Could not fetch question.", nil, http.StatusInternalServerError, w)
			fmt.Println(err)
		}
		return
	}

	ResponseWriter(true, "Hints fetched.", hints, http.StatusOK, w)
}

func (hawk *App) getStats(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)
	/*
		if (currUser == User{}) {
			ResponseWriter(false, "User not logged in.", nil, http.StatusNetworkAuthenticationRequired, w)
			return
		}
	*/
	currStats := Stats{}
	//Total players
	err := hawk.DB.Model(&User{}).Count(&currStats.TotalPlayers).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cant get total players", nil, http.StatusInternalServerError, w)
		return
	}

	//Total answer attempts
	err = hawk.DB.Model(&Attempt{}).Count(&currStats.AnswerAttempts).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cant get total answer attempts", nil, http.StatusInternalServerError, w)
		return
	}
	/*
		//on same level in linear gameplay
		err = hawk.DB.Model(&User{}).Where("region5 = ?", currUser.Region5).Count(&currStats.SameLevel).Error
		if err != nil {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Cant get players at par", nil, http.StatusInternalServerError, w)
			return
		}
		//leading on linear gameplay
		err = hawk.DB.Model(&User{}).Where("region5 > ?", currUser.Region5).Count(&currStats.Leading).Error
		if err != nil {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Cant get leading players at par", nil, http.StatusInternalServerError, w)
			return
		}

	*/
	//same points
	err = hawk.DB.Model(&User{}).Where("points = ?", currUser.Points).Count(&currStats.SameLevel).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cant get players at par", nil, http.StatusInternalServerError, w)
		return
	}
	//leading
	err = hawk.DB.Model(&User{}).Where("points > ?", currUser.Points).Count(&currStats.Leading).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cant get leading users", nil, http.StatusInternalServerError, w)
		return
	}

	
	currStats.Trailing = currStats.TotalPlayers - (currStats.Leading + currStats.SameLevel)
	currStats.SameLevel -= 1 //counts currUser also
	ResponseWriter(true, "Current stats", currStats, http.StatusOK, w)

}
