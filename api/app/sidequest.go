package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

const (
	PointsPerQuestion       = 1
	UnlockRegionPoints      = 3
	TotalSidequestQuestions = 6
)

func (hawk *App) getSidequestQuestion(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)
	currLevel := currUser.Region0
	if currLevel == TotalSidequestQuestions+1 {
		ResponseWriter(true, "Sidequest over", nil, http.StatusOK, w)
		return
	}
	index := (currLevel - 1) * 2
	level := currUser.SidequestOrder[index] - 48 //ascii to int
	question := Question{}
	err := hawk.DB.Select("id, level, region, question, add_info").Where("region = ? AND level = ?", 0, level).First(&question).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Question does not exist", nil, http.StatusNotFound, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}

	}
	ResponseWriter(true, "Question fetched", question, http.StatusOK, w)
}

func (hawk *App) checkSidequestAnswer(w http.ResponseWriter, r *http.Request) {

	currUser := r.Context().Value("User").(User)
	if currUser.IsVerified != 1 {
		ResponseWriter(false, "Not verified", nil, http.StatusOK, w)
		return
	}
	if currUser.Banned == 1 {
		ResponseWriter(false, "User banned", nil, http.StatusOK, w)
		return
	}
	//obtain answer from request, region is Region0 by default for sidequest
	checkAns := CheckAnswer{}
	err := json.NewDecoder(r.Body).Decode(&checkAns)
	if err != nil {
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusBadRequest, w)
		return
	} //sanitize answer
	checkAns.Answer = Sanitize(checkAns.Answer)
	//actual level
	actualLevel := currUser.SidequestOrder[(currUser.Region0-1)*2] - 48
	//find actual answer
	question := Question{}
	err = hawk.DB.Where("level = ? AND region = ?", actualLevel, 0).First(&question).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Question does not exist", nil, http.StatusNotFound, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//Check if equal, close or wrong
	status := CheckAnswerStatus(checkAns.Answer, question.Answer)
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
	tx.Commit()
	if status == CorrectAnswer {
		//update level and sidequest points
		currUser.Region0 += 1
		currUser.SideQuestPoints += PointsPerQuestion
		tx = hawk.DB.Begin()
		err = tx.Model(&currUser).Updates(map[string]interface{}{"region0": currUser.Region0, "side_quest_points": currUser.SideQuestPoints}).Error
		/*
			err = tx.Update("region0", currUser.Region0).Error

			if err != nil {
				LogRequest(r, ERROR, err.Error())
				ResponseWriter(false, "Could not update database", nil, http.StatusInternalServerError, w)
				tx.Rollback()
				return
			}
			err = tx.Update("side_quest_points", currUser.SideQuestPoints).Error
			//err = tx.Model(&currUser).Updates(map[string]interface{}{"Region0": currUser.Region0+1, "side_quest_points": currUser.SideQuestPoints+PointsPerQuestion}).Error

			if err != nil {
				LogRequest(r, ERROR, err.Error())
				ResponseWriter(false, "Could not update database", nil, http.StatusInternalServerError, w)
				tx.Rollback()
				return
			}
		*/
		tx.Commit()
	}
	ResponseWriter(true, "Answer status", status, http.StatusOK, w)
	return
}

func (hawk *App) unlockRegion(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)

	if currUser.SideQuestPoints < UnlockRegionPoints {
		ResponseWriter(false, "Not enough points to unlock region", nil, http.StatusForbidden, w)
		return
	}

	var err error
	tx := hawk.DB.Begin()

	nextRegion := GetNextRegion(currUser.UnlockOrder)
	if nextRegion == '5' {
		ResponseWriter(false, "Cannot unlock linear region from sidequest", nil, http.StatusOK, w)
		return
	}
	currUser.UnlockOrder = UpdateUnlockOrder(currUser.UnlockOrder, int(nextRegion)-48)

	err = tx.Model(&currUser).Update("SideQuestPoints", currUser.SideQuestPoints-UnlockRegionPoints).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}

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
		LogRequest(r, ERROR, err.Error())
		fmt.Println("Could not unlock region in DB")
		ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}

	err = tx.Model(&currUser).Update("unlock_order", currUser.UnlockOrder).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		fmt.Println("Could not unlock region in DB")
		ResponseWriter(false, "Could not unlock region in DB", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	ResponseWriter(true, "Region unlocked", nil, http.StatusOK, w)
}
