package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CheckAnswer struct {
	Answer	string
	Regionid	int
	Level 		int
}

func (hawk *App) checkAnswer (w http.ResponseWriter, r *http.Request){
	//obtain answerid, regionid, level, answer from request
	checkAns := CheckAnswer {}
	err := json.NewDecoder(r.Body).Decode (&checkAns)
	if  err != nil {
		fmt.Println ("Could not decode checkAnswer struct")
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusInternalServerError, w)
		return
	}
	//sanitize answer
	checkAns.Answer = strings.TrimSpace(checkAns.Answer)
	checkAns.Answer = strings.ToLower(checkAns.Answer)
	//find actual answer
	question := Question {}
	err = hawk.DB.Where ("level = ? AND region = ?", checkAns.Level, checkAns.Regionid).First (&question).Error
	actualAns := question.Answer

	//check if same, close or wrong
	status := checkAnswerStatus (checkAns.Answer, actualAns)


	ResponseWriter(true,"Answer status", status,http.StatusOK, w)

}

func (hawk *App) getQuestion (w http.ResponseWriter, r *http.Request) {
	keys ,ok := r.URL.Query()["region"]
	if !ok || len(keys[0])<1 {
		ResponseWriter(false,"Invalid request",nil,http.StatusBadRequest, w)
		return
	}

	currUser := r.Context().Value("CurrUser").(CurrUser)

	key := keys[0]
	level := 0

	switch key {
	case "1": level = currUser.Region1
	case "2": level = currUser.Region2
	case "3": level = currUser.Region3
	case "4": level = currUser.Region4
	case "5": level = currUser.Region5
	}

	question := Question{}

	err := hawk.DB.Where("level=? AND region=?", level, key).First(&question).Error
	if err != nil {
		ResponseWriter(false,"Could not fetch question.",nil,http.StatusInternalServerError, w)
		return
	}

	ResponseWriter(true,"Question fetched.", question, http.StatusOK, w)
}

func (hawk *App) getHints (w http.ResponseWriter, r *http.Request) {
	keys ,ok := r.URL.Query()["question"]
	if !ok || len(keys[0])<1 {
		ResponseWriter(false,"Invalid request",nil,http.StatusBadRequest, w)
		return
	}

	key := keys[0]

	var hints []string

	err := hawk.DB.Model(&Hint{}).Where("question=? AND active=1", key).Pluck("hint",&hints).Error
	if err != nil {
		ResponseWriter(false,"Could not fetch hint.",hints,http.StatusInternalServerError, w)
		return
		}

	ResponseWriter(true, "Hints fetched.", hints, http.StatusOK, w)
}
