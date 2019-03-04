package app

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (hawk *App) addQuestion(w http.ResponseWriter, r *http.Request) {
	newQues := Question{}

	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		ResponseWriter(false, "Could not parse request body", nil, http.StatusInternalServerError, w)
		return
	}

	newQues.Question = strings.TrimSpace(newQues.Question)
	newQues.AddInfo = strings.TrimSpace(newQues.AddInfo)
	newQues.AddedBy = strings.TrimSpace(newQues.AddedBy)
	newQues.Answer = Sanitize(newQues.Answer)

	tx := hawk.DB.Begin()
	err = tx.Create(&newQues).Error

	if err != nil {
		ResponseWriter(false, "Database error", err.Error(), http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}

	tx.Commit()
	ResponseWriter(true, "Question added", nil, http.StatusOK, w)
}

func (hawk *App) addHint(w http.ResponseWriter, r *http.Request) {
	newHint := Hint{}

	err := json.NewDecoder(r.Body).Decode(&newHint)
	if err != nil {
		ResponseWriter(false, "Could not parse request body", nil, http.StatusInternalServerError, w)
		return
	}

	newHint.Hint = strings.TrimSpace(newHint.Hint)
	newHint.Active = 0

	tx := hawk.DB.Begin()
	err = tx.Create(&newHint).Error

	if err != nil {
		ResponseWriter(false, "Database error", err.Error(), http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint added", nil, http.StatusOK, w)
}
