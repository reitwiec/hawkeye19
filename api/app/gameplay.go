package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func (hawk *App) checkAnswer(w http.ResponseWriter, r *http.Request) {
	//obtain currUser from context
	currUser := r.Context().Value("CurrUser").(CurrUser)
	//obtain answer, regionId, level from request
	checkAns := CheckAnswer{}
	err := json.NewDecoder(r.Body).Decode(&checkAns)
	if err != nil {
		fmt.Println("Could not decode checkAnswer struct " + err.Error())
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusBadRequest, w)
		return
	}


	//sanitize answer
	checkAns.Answer = strings.TrimSpace(checkAns.Answer)
	checkAns.Answer = strings.ToLower(checkAns.Answer)

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
		return
	}
	//if answer is correct
	if status == CorrectAnswer {
		//answer is correct
		//update level in region
		switch checkAns.RegionId {
		case 1:
			currUser.Region1 += 1
		case 2:
			currUser.Region2 += 1
		case 3:
			currUser.Region3 += 1
		case 4:
			currUser.Region4 += 1
		case 5:
			currUser.Region5 += 1
		}

		user := User{}
		hawk.DB.Where("ID = ?", currUser.ID).First(&user)
		tx := hawk.DB.Begin()
		region := "Region" + strconv.Itoa(checkAns.RegionId)
		err = tx.Model(&user).Update(region, checkAns.Level+1).Error
		if err != nil {
			fmt.Println("Could not update level")
			ResponseWriter(false, "Could not update level", nil, http.StatusInternalServerError, w)
			tx.Rollback()
		}
		//set new cookie
		err = SetSession(w, currUser, 86400)
		if err != nil {
			fmt.Println("Could not set cookie")
			ResponseWriter(false, "Could not set cookie in checking answer", nil, http.StatusInternalServerError, w)
			tx.Rollback()
		}
		//commit transaction
		tx.Commit()
	}
	ResponseWriter(true, "Answer status", status, http.StatusOK, w)
}

func (hawk *App) getQuestion(w http.ResponseWriter, r *http.Request) {

	currUser := r.Context().Value("CurrUser").(CurrUser)
	//@TODO: Implement in middleware
	if (currUser == CurrUser{}) {
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

	err := hawk.DB.Where("level=? AND region=?", level, key).First(&question).Error
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
		ResponseWriter(false, "Could not fetch hint.", hints, http.StatusInternalServerError, w)
		return
	}

	ResponseWriter(true, "Hints fetched.", hints, http.StatusOK, w)
}

func (hawk *App) getRecentTries (w http.ResponseWriter, r *http.Request){
	currUser := r.Context().Value ("CurrUser").(CurrUser)
	if currUser == (CurrUser {}){
		fmt.Println ("Not logged in")
		ResponseWriter(false, "Not logged in", nil, http.StatusNetworkAuthenticationRequired, w)
		return
	}
	//get user ID from context, region and level from GET request
	keys, _ := r.URL.Query ()["question"]
	fmt.Println (keys[0])
}

func (hawk *App) getStats(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("CurrUser").(CurrUser)
	if (currUser == CurrUser{}) {
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
