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
	err := hawk.DB.Where ("level = ? AND region = ?", checkAns.Level, checkAns.Regionid).First (&question).Error
	actualAns := question.Answer

	//check if same, close or wrong
	status := checkAnswerStatus (checkAns.Answer, actualAns)

}
