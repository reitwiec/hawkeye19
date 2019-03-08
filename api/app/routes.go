package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (hawk *App) LoadRoutes() {

	hawk.router = mux.NewRouter()
	//Auth routes
	hawk.router.HandleFunc("/api/addUser", hawk.addUser).Methods("POST")
	hawk.router.HandleFunc("/api/login", hawk.login).Methods("POST")
	hawk.router.HandleFunc("/api/logout", hawk.logout).Methods("POST")
	hawk.router.HandleFunc("/api/forgotPassword", hawk.forgotPassword).Methods("POST")
	hawk.router.HandleFunc("/api/resetPassword", hawk.resetPassword).Methods("POST")
	hawk.router.HandleFunc("/api/checkUsername", hawk.checkUsername).Methods("POST")
	hawk.router.HandleFunc("/api/checkEmail", hawk.checkEmail).Methods("POST")

	//Gameplay routes
	hawk.router.HandleFunc("/api/checkAnswer", hawk.createContext(hawk.checkAnswer, false)).Methods("POST")
	hawk.router.HandleFunc("/api/getQuestion", hawk.createContext(hawk.getQuestion, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getHints", hawk.createContext(hawk.getHints, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getStats", hawk.createContext(hawk.getStats, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getRecentTries", hawk.createContext(hawk.getRecentTries, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getSideQuestQuestion", hawk.createContext(hawk.getSideQuestQuestion, false)).Methods("GET")
	hawk.router.HandleFunc("/api/checkAnswer", hawk.createContext(hawk.checkAnswer, false)).Methods("POST")
	hawk.router.HandleFunc("/api/getQuestion", hawk.createContext(hawk.getQuestion, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getHints", hawk.createContext(hawk.getHints, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getStats", hawk.createContext(hawk.getStats, false)).Methods("GET")
	hawk.router.HandleFunc("/api/getRecentTries", hawk.createContext(hawk.getRecentTries, false)).Methods("GET")

	//Admin gameplay routes
	hawk.router.HandleFunc("/api/addQuestion", hawk.createContext(hawk.addQuestion, true)).Methods("POST")
	hawk.router.HandleFunc("/api/addHint", hawk.createContext(hawk.addHint, true)).Methods("POST")
	hawk.router.HandleFunc("/api/editQuestion", hawk.createContext(hawk.editQuestion, true)).Methods("POST")
	hawk.router.HandleFunc("/api/editHint", hawk.createContext(hawk.editHint, true)).Methods("POST")
	hawk.router.HandleFunc("/api/deleteHint", hawk.createContext(hawk.deleteHint, true)).Methods("GET")
	hawk.router.HandleFunc("/api/deleteQuestion", hawk.createContext(hawk.deleteQuestion, true)).Methods("GET")
	hawk.router.HandleFunc("/api/activateHint", hawk.createContext(hawk.activateHint, true)).Methods("PUT")
	hawk.router.HandleFunc("/api/deactivateHint", hawk.createContext(hawk.deactivateHint, true)).Methods("PUT")
	hawk.router.HandleFunc("/api/listQuestions", hawk.createContext(hawk.listQuestions, true)).Methods("GET")
	hawk.router.HandleFunc("/api/listHints", hawk.createContext(hawk.listHints, true)).Methods("GET")

	//Admin User routes
	hawk.router.HandleFunc("/api/editUser", hawk.createContext(hawk.editUser, true)).Methods("POST")
	hawk.router.HandleFunc("/api/deleteUser", hawk.createContext(hawk.deleteUser, true)).Methods("GET")
	hawk.router.HandleFunc("/api/banUser", hawk.createContext(hawk.banUser, true)).Methods("GET")
	hawk.router.HandleFunc("/api/unbanUser", hawk.createContext(hawk.unbanUser, true)).Methods("GET")
	hawk.router.HandleFunc("/api/makeAdmin", hawk.createContext(hawk.makeAdmin, true)).Methods("GET")
	hawk.router.HandleFunc("/api/revokeAdmin", hawk.createContext(hawk.revokeAdmin, true)).Methods("GET")
	hawk.router.HandleFunc("/api/listUsers", hawk.createContext(hawk.listUsers, true)).Methods("GET")
	hawk.router.HandleFunc("/api/searchUser", hawk.createContext(hawk.searchUser, true)).Methods("GET")
	hawk.router.HandleFunc("/api/userLogs", hawk.createContext(hawk.userLogs, true)).Methods("GET")
}

func (hawk *App) createContext(next http.HandlerFunc, isAdmin bool) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//extract data from cookie
			currUser, err := GetCurrUser(w, r)
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
			next.ServeHTTP(w, r)
		})
}
