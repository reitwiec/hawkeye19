package app

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (hawk *App) LoadRoutes() {

	hawk.router = mux.NewRouter()
	//Auth routes
	hawk.router.HandleFunc("/api/addUser", hawk.createContext(hawk.addUser, false, false)).Methods("POST")
	hawk.router.HandleFunc("/api/login", hawk.createContext(hawk.login, false, false)).Methods("POST")
	hawk.router.HandleFunc("/api/logout", hawk.createContext(hawk.logout, false, false)).Methods("GET")
	hawk.router.HandleFunc("/api/forgotPassword", hawk.createContext(hawk.forgotPassword, false, false)).Methods("POST")
	hawk.router.HandleFunc("/api/resetPassword", hawk.createContext(hawk.resetPassword, false, false)).Methods("POST")
	hawk.router.HandleFunc("/api/verifyUser", hawk.createContext(hawk.verifyUser, false, false)).Methods("GET")

	//Gameplay routes
	hawk.router.HandleFunc("/api/checkAnswer", hawk.createContext(hawk.checkAnswer, false, true)).Methods("POST")
	hawk.router.HandleFunc("/api/getQuestion", hawk.createContext(hawk.getQuestion, false, true)).Methods("GET")
	hawk.router.HandleFunc("/api/getHints", hawk.createContext(hawk.getHints, false, true)).Methods("GET")
	hawk.router.HandleFunc("/api/getStats", hawk.createContext(hawk.getStats, false, true)).Methods("GET")
	hawk.router.HandleFunc("/api/getRecentTries", hawk.createContext(hawk.getRecentTries, false, true)).Methods("GET")

	//Help center routes
	hawk.router.HandleFunc("/api/submitHelpRequest", hawk.createContext(hawk.submitHelpRequest, false, true)).Methods("POST")
	hawk.router.HandleFunc("/api/getHelpRequests", hawk.createContext(hawk.getHelpRequests, false, true)).Methods("GET")

	//Sidequest routes
	hawk.router.HandleFunc("/api/getSidequestQuestion", hawk.createContext(hawk.getSidequestQuestion, false, true)).Methods("GET")
	hawk.router.HandleFunc("/api/checkSidequestAnswer", hawk.createContext(hawk.checkSidequestAnswer, false, true)).Methods("POST")
	hawk.router.HandleFunc("/api/unlockRegion", hawk.createContext(hawk.unlockRegion, false, true)).Methods("GET")

	//Admin gameplay routes
	hawk.router.HandleFunc("/api/addQuestion", hawk.createContext(hawk.addQuestion, true, true)).Methods("POST")
	hawk.router.HandleFunc("/api/addHint", hawk.createContext(hawk.addHint, true, true)).Methods("POST")
	hawk.router.HandleFunc("/api/editQuestion", hawk.createContext(hawk.editQuestion, true, true)).Methods("POST")
	hawk.router.HandleFunc("/api/editHint", hawk.createContext(hawk.editHint, true, true)).Methods("POST")
	hawk.router.HandleFunc("/api/deleteHint", hawk.createContext(hawk.deleteHint, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/deleteQuestion", hawk.createContext(hawk.deleteQuestion, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/activateHint", hawk.createContext(hawk.activateHint, true, true)).Methods("PUT")
	hawk.router.HandleFunc("/api/deactivateHint", hawk.createContext(hawk.deactivateHint, true, true)).Methods("PUT")
	hawk.router.HandleFunc("/api/listQuestions", hawk.createContext(hawk.listQuestions, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/listHints", hawk.createContext(hawk.listHints, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/questionLogs", hawk.createContext(hawk.questionLogs, true, true)).Methods("GET")

	//Admin User routes
	hawk.router.HandleFunc("/api/editUser", hawk.createContext(hawk.editUser, true, true)).Methods("POST")
	hawk.router.HandleFunc("/api/deleteUser", hawk.createContext(hawk.deleteUser, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/banUser", hawk.createContext(hawk.banUser, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/unbanUser", hawk.createContext(hawk.unbanUser, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/makeAdmin", hawk.createContext(hawk.makeAdmin, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/revokeAdmin", hawk.createContext(hawk.revokeAdmin, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/listUsers", hawk.createContext(hawk.listUsers, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/searchUser", hawk.createContext(hawk.searchUser, true, true)).Methods("GET")
	hawk.router.HandleFunc("/api/userLogs", hawk.createContext(hawk.userLogs, true, true)).Methods("GET")

	//Helper routes
	hawk.router.HandleFunc("/api/getUser", hawk.GetUser).Methods("GET")

}

func (hawk *App) createContext(next http.HandlerFunc, isAdmin bool, isLoggedIn bool) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if isLoggedIn {
				//extract data from cookie
				currUser, err := GetCurrUser(r)
				if err != nil {
					fmt.Println("Not logged in")
					ResponseWriter(false, "Not logged in", nil, http.StatusForbidden, w)
					return
				}
				user := User{}
				//gets rest of the data from db
				err = hawk.DB.First(&user, currUser.ID).Error
				if err != nil {
					fmt.Println("Database error " + err.Error())
					return
				}
				user.Password = ""
				//for only admins
				if isAdmin {
					//check if Admin
					if user.Access != 1 {
						fmt.Println("Not an admin")
						ResponseWriter(false, "Not an admin", user, http.StatusForbidden, w)
						return
					}
				}
				//create new context with CurrUser
				ctx := context.WithValue(r.Context(), "User", user)
				r = r.WithContext(ctx)
			}
			LogRequest(r, INFO, "")
			next.ServeHTTP(w, r)
		})
}
