package app

import (
	"encoding/json"
	"net/http"
)

func (hawk *App) addQuestion (w http.ResponseWriter, r http.Request) {
	newQues := Question{}

	err := json.NewDecoder(r.Body).Decode(&newQues)
	if err != nil {
		ResponseWriter(false,"Could not parse request body",nil, http.StatusInternalServerError, w)
		return
	}

	
}