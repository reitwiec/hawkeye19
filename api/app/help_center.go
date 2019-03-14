package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const requestsPerPage = 15

func (hawk *App) submitHelpRequest(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("User").(User)

	newHelp := Help{}

	err := json.NewDecoder(r.Body).Decode(&newHelp)
	if err != nil {
		fmt.Println("Could not decode checkAnswer struct " + err.Error())
		ResponseWriter(false, "Could not decode check answer struct", nil, http.StatusBadRequest, w)
		return
	}
	
	newHelp.User = currUser.ID
	newHelp.Username = currUser.Username
	newHelp.Email = currUser.Email
	newHelp.Timestamp = time.Now()

	tx := hawk.DB.Begin()

	err = tx.Create(&newHelp).Error
	if err != nil {
		ResponseWriter(false, "Could not add help request.", nil, http.StatusInternalServerError, w)
		return
	}

	tx.Commit()
	ResponseWriter(true, "Help request added.", nil, http.StatusOK, w)
}

func (hawk *App) getHelpRequests (w http.ResponseWriter, r *http.Request) {
	pgs, ok := r.URL.Query()["page"]
	if !ok || len(pgs[0]) < 1 {
		ResponseWriter(false, "Cant list help requests.", nil, http.StatusBadRequest, w)
		return
	}
	page, err := strconv.Atoi(pgs[0])
	if err != nil {
		ResponseWriter(false, "Parameters not valid. Cannot list help requests.", nil, http.StatusBadRequest, w)
		return
	}
	var helpRequests []Help

	offset := (page - 1) * requestsPerPage
	err = hawk.DB.Find(&helpRequests).Offset(offset).Limit(requestsPerPage).Error
	if err != nil {
		ResponseWriter(false, "Cannot list questions.", nil, http.StatusInternalServerError, w)
		return
	}

	ResponseWriter(true, "List of help requests.", helpRequests, http.StatusOK, w)
}
