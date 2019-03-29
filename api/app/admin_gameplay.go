package app

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const quesPerPage = 15
const quesLogsPerPage = 15

func (hawk *App) addQuestion(w http.ResponseWriter, r *http.Request) {
	newQues := Question{}

	currUser := r.Context().Value("User").(User)

	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not parse request body", nil, http.StatusInternalServerError, w)
		return
	}

	newQues.Question = strings.TrimSpace(newQues.Question)
	newQues.AddInfo = strings.TrimSpace(newQues.AddInfo)
	newQues.AddedBy = currUser.Username
	newQues.Answer = SanitizeAnswer(newQues.Answer)
	newQues.Timestamp = time.Now()

	tx := hawk.DB.Begin()
	err = tx.Create(&newQues).Error

	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}

	tx.Commit()
	ResponseWriter(true, "Question added", nil, http.StatusOK, w)
}

func (hawk *App) addHint(w http.ResponseWriter, r *http.Request) {
	newHint := Hint{}
	currUser := r.Context().Value("User").(User)

	err := json.NewDecoder(r.Body).Decode(&newHint)
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not parse request body", nil, http.StatusInternalServerError, w)
		return
	}

	newHint.Hint = strings.TrimSpace(newHint.Hint)
	newHint.Active = 0
	newHint.AddedBy = currUser.Username
	newHint.Timestamp = time.Now()
	tx := hawk.DB.Begin()
	err = tx.Create(&newHint).Error

	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint added", newHint.ID, http.StatusOK, w)
}

func (hawk *App) editQuestion(w http.ResponseWriter, r *http.Request) {
	newQues := Question{}
	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		ResponseWriter(false, "Question not updated", nil, http.StatusBadRequest, w)
		return
	}

	newQues.Question = strings.TrimSpace(newQues.Question)
	newQues.Answer = Sanitize(newQues.Answer)
	newQues.AddInfo = strings.TrimSpace(newQues.AddInfo)

	q := Question{}
	tx := hawk.DB.Begin()
	err = tx.Where("level = ? and region = ?", newQues.Level, newQues.Region).First(&q).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Question not found", nil, http.StatusBadRequest, w)
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Question not updated", nil, http.StatusInternalServerError, w)
		}
		return
	}
	err = tx.Model(&q).Updates(Question{Question: newQues.Question, Answer: newQues.Answer, AddInfo: newQues.AddInfo}).Error
	if err != nil {
		tx.Rollback()
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Hint not deactivated", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "Hint deactivated", nil, http.StatusOK, w)
}

func (hawk *App) listQuestions(w http.ResponseWriter, r *http.Request) {
	pgs, ok := r.URL.Query()["page"]
	if !ok || len(pgs[0]) < 1 {
		ResponseWriter(false, "Cant list questions", nil, http.StatusBadRequest, w)
		return
	}
	page, err := strconv.Atoi(pgs[0])
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Parameters not valid. Cannot list questions.", nil, http.StatusBadRequest, w)
		return
	}
	var questions []Question
	offset := (page - 1) * quesPerPage
	err = hawk.DB.Find(&questions).Offset(offset).Order("level asc").Limit(quesPerPage).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cannot list questions", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "List of questions", questions, http.StatusOK, w)

}

func (hawk *App) listHints(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Cant list hints", nil, http.StatusBadRequest, w)
		return
	}
	id, err := strconv.Atoi(keys[0])
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Parameters not valid Cannot display hints.", nil, http.StatusBadRequest, w)
		return
	}
	var hints []Hint
	err = hawk.DB.Where(" id = ?", id).Find(&hints).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cannot list hints", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "List of hints", hints, http.StatusOK, w)
}

func (hawk *App) questionLogs(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	if len(keys) < 2 {
		ResponseWriter(false, "Cannot log answers.", nil, http.StatusBadRequest, w)
		return
	}
	page, err := strconv.Atoi(keys["page"][0])
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Parameters not valid.Cannot list answers", nil, http.StatusBadRequest, w)
		return
	}
	id, err := strconv.Atoi(keys["question"][0])
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Parameters not valid. Cannot list answers.", nil, http.StatusBadRequest, w)
		return
	}
	var answers []Attempt
	offset := (page - 1) * perPage
	err = hawk.DB.Where("question = ?", id).Order("timestamp desc").Offset(offset).Limit(perPage).Find(&answers).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Cannot log answers", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "Requested answer logs", answers, http.StatusOK, w)
}
