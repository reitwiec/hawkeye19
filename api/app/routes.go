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
	hawk.router.HandleFunc("/api/addUser", hawk.addUser).Methods("POST")
	hawk.router.HandleFunc("/api/login", hawk.login).Methods("POST")
	hawk.router.HandleFunc("/api/logout", hawk.logout).Methods("POST")
	hawk.router.HandleFunc("/api/forgotPassword", hawk.forgotPassword).Methods("POST")
	hawk.router.HandleFunc("/api/resetPassword", hawk.resetPassword).Methods("POST")
	hawk.router.HandleFunc("/api/checkUsername", hawk.checkUsername).Methods("POST")
	hawk.router.HandleFunc("/api/checkEmail", hawk.checkEmail).Methods("POST")

	//Gameplay routes

	hawk.router.HandleFunc("/api/checkAnswer", hawk.createContext(hawk.checkAnswer)).Methods("POST")
	hawk.router.HandleFunc("/api/getQuestion", hawk.createContext(hawk.getQuestion)).Methods("GET")
	hawk.router.HandleFunc("/api/getHints", hawk.createContext(hawk.getHints)).Methods("GET")
	hawk.router.HandleFunc("/api/getStats", hawk.createContext(hawk.getStats)).Methods("GET")
	hawk.router.HandleFunc("/api/getRecentTries", hawk.createContext(hawk.getRecentTries)).Methods("GET")
}

func (hawk *App) createContext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//extract data from cookie
			currUser, err := GetCurrUser(w, r)
			if err != nil {
				fmt.Println("Could not read cookie data")
				ResponseWriter(false, "Could not read cookie data", nil, http.StatusInternalServerError, w)
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
			//create new context with CurrUser
			ctx := context.WithValue(r.Context(), "User", user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}

/*
//sample method to see if context is working
func (hawk *App) checkContext(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("CurrUser").(CurrUser)
	fmt.Println(currUser)
}
*/
