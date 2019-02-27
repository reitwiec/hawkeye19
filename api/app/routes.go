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
	hawk.router.HandleFunc("/api/checkContext", hawk.createContext(hawk.checkContext)).Methods("POST")

	//Gameplay routes
	hawk.router.HandleFunc("/api/checkAnswer",hawk.createContext(hawk.checkAnswer)).Methods("POST")
	hawk.router.HandleFunc("/api/getQuestion",hawk.createContext(hawk.getQuestion)).Methods("GET")
}

func (hawk *App) createContext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//extract data from cookie
			currUser := GetCurrUser(w, r)
			if currUser == (CurrUser{}) {
				fmt.Println("Could not read cookie data")
				return
			}
			user := User{}
			//gets rest of the data from db
			err := hawk.DB.First(&user, currUser.ID).Error
			if err != nil {
				fmt.Println("Database error " + err.Error())
				return
			}
			currUser.Points = user.Points
			currUser.Access = user.Access
			currUser.Region1 = user.Region1
			currUser.Region2 = user.Region2
			currUser.Region3 = user.Region3
			currUser.Region4 = user.Region4
			currUser.Region5 = user.Region5
			//create new context with CurrUser
			ctx := context.WithValue(r.Context(), "CurrUser", currUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}

//sample method to see if context is working
func (hawk *App) checkContext(w http.ResponseWriter, r *http.Request) {
	currUser := r.Context().Value("CurrUser").(CurrUser)
	fmt.Println(currUser)
}