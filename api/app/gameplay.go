package app

import (
	"net/http"
	//"encoding/json"
	//"fmt"
	//"github.com/jinzhu/gorm"
)

type InitGameplay struct {
	Username string `json:"username"`
	Region1 int `json:"region1"`
	Region2 int `json:"region1"`
	Region3 int `json:"region1"`
	Region4 int `json:"region1"`
	Region5 int `json:"region1"`
}

func (hawk *App) intiGameplay(w http.ResponseWriter, r *http.Request) {
	currUser, err := GetCurrentUser(w,r)
	if err!=nil {
		ResponseWriter(false,"Not logged in",nil,http.StatusForbidden,w)
		return
	}

	userData := InitGameplay{}
	userData.Username = currUser.Username
	userData.Region1 = currUser.Region1
	userData.Region2 = currUser.Region2
	userData.Region3 = currUser.Region3
	userData.Region4 = currUser.Region4
	userData.Region5 = currUser.Region5

	ResponseWriter(true, "Region detais found", userData, http.StatusOK, w)
	return
}