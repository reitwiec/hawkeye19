package app

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func (hawk *App) addQuestion(w http.ResponseWriter, r *http.Request) {
	newQues := Question{}

	currUser := r.Context().Value("User").(User)
	if (currUser == User{}) {
		ResponseWriter(false, "User not logged in.", nil, http.StatusNetworkAuthenticationRequired, w)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		ResponseWriter(false, "Could not parse request body", nil, http.StatusInternalServerError, w)
		return
	}

	newQues.Question = strings.TrimSpace(newQues.Question)
	newQues.AddInfo = strings.TrimSpace(newQues.AddInfo)
	newQues.AddedBy = currUser.Username
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

func (hawk *App) editQuestion(w http.ResponseWriter, r *http.Request) {
	newQues := Question{}
	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		ResponseWriter(false, "Question not updated", nil, http.StatusBadRequest, w)
		return
	}

	newQues.Question = strings.TrimSpace(newQues.Question)
	newQues.Answer = strings.TrimSpace(newQues.Answer)
	newQues.AddInfo = strings.TrimSpace(newQues.AddInfo)

	q := Question{}
	tx := hawk.DB.Begin()
	err = tx.Where("level = ? and region=?", newQues.Level,newQues.Region).First(&q).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Question not found", nil, http.StatusBadRequest, w)
		} else {
			ResponseWriter(false, "Question not updated", nil, http.StatusInternalServerError, w)
		}
		return
	}
	err = tx.Model(&q).Updates(Question{Question: newQues.Question, Answer: newQues.Answer, AddInfo: newQues.AddInfo}).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Question not updated", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Question edited", nil, http.StatusOK, w)
}

func (hawk *App) editHint(w http.ResponseWriter, r *http.Request) {
	newHint := Hint{}
	err := json.NewDecoder(r.Body).Decode(&newHint)
	if err != nil {
		ResponseWriter(false, "Hint not updated", nil, http.StatusBadRequest, w)
		return
	}
	newHint.Hint = strings.TrimSpace(newHint.Hint)

	h := Hint{}
	tx := hawk.DB.Begin()
	err = tx.Model(&Hint{}).Where("id = ?", newHint.ID).First(&h).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Hint not found", nil, http.StatusBadRequest, w)
		} else {
			ResponseWriter(false, "Hint not updated", nil, http.StatusBadRequest, w)
		}
		return
	}
	err = tx.Model(&h).Updates(Hint{Hint: newHint.Hint, Question: newHint.Question, Active: newHint.Active}).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Hint not updated", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint edited", nil, http.StatusOK, w)
}

func (hawk *App) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Question not deleted.", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]

	tx := hawk.DB.Begin()
	q := Question{}
	err := tx.Where("id=?", key).First(&q).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Question not found", nil, http.StatusBadRequest, w)
		} else {
			tx.Rollback()
			ResponseWriter(false, "Question not deleted", nil, http.StatusInternalServerError, w)
		}
		return
	}
	err = tx.Delete(q).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Question not deleted", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Question deleted.", nil, http.StatusOK, w)
}

func (hawk *App) deleteHint(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Hint not deleted.", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]

	tx := hawk.DB.Begin()
	h := Hint{}
	err := tx.Where("id = ?", key).First(&h).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Hint not found", nil, http.StatusBadRequest, w)
		} else {
			tx.Rollback()
			ResponseWriter(false, "Hint not deleted", nil, http.StatusInternalServerError, w)
		}
		return
	}
	err = tx.Delete(h).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Hint not deleted", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint deleted", nil, http.StatusOK, w)
}

func (hawk *App) activateHint(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Hint not activated", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]

	tx := hawk.DB.Begin()
	h := Hint{}
	err := tx.Where("ID = ?", key).First(&h).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Hint not found", nil, http.StatusBadRequest, w)
		} else {
			tx.Rollback()
			ResponseWriter(false, "Hint not activated", nil, http.StatusInternalServerError, w)
		}
		return
	}

	err = tx.Model(h).Update("Active", 1).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Hint not activated", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint activated", nil, http.StatusOK, w)
}

func (hawk *App) deactivateHint(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Hint not deactivated", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]
	h := Hint{}
	tx := hawk.DB.Begin()
	err := tx.Where("ID = ?", key).First(&h).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Hint not found", nil, http.StatusBadRequest, w)
		} else {
			tx.Rollback()
			ResponseWriter(false, "Hint not deactivated", nil, http.StatusInternalServerError, w)
		}
		return
	}

	err = tx.Model(h).Update("Active", 0).Error
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Hint not deactivated", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint deactivated", nil, http.StatusOK, w)
}