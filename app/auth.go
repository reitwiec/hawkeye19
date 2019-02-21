package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strings"
)

func (hawk * App) addUser (w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Decoding error, user not registered")
		ResponseWriter(false, "Bad Request", nil, http.StatusBadRequest, w)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			ResponseWriter(false, "Invalid validation error ", nil, http.StatusBadRequest, w)
			return
		}
		var fields [] string
		for _, err := range err.(validator.ValidationErrors) {
			fields = append(fields, err.Field())
		}
		ResponseWriter(false, "Bad Request", fields, http.StatusBadRequest, w)
		return
	}

	//validation successful
	//hash  and salt password
	hash, err := bcrypt.GenerateFromPassword([]byte (user.Password), 14)
	if err != nil {
		log.Println(err)
		ResponseWriter(false, "Error in hash and salt", nil, http.StatusInternalServerError, w)
		return
	}

	newUser := User{
		Username: strings.TrimSpace(user.Username),
		Name:     strings.TrimSpace(user.Name),
		Password: string(hash),
		Access:   0,
		Email:    strings.TrimSpace(user.Email),
		Tel:      strings.TrimSpace(user.Tel),
		College:  strings.TrimSpace(user.College),
		Level:    1,
		Banned:   0,
		Points:   2,
	}
	//load newUser to database
	//hawk.DB.CreateTable (&User{})
	//begin transaction
	tx := hawk.DB.Begin ()
	err = tx.Create (&newUser).Error
	if err != nil {
		fmt.Println ("Database error, new user not created")
		ResponseWriter(false, "Database error, new user not created", nil, http.StatusInternalServerError, w)
		tx.Rollback ()
		return
	}
	tx.Commit ()
	fmt.Println ("New user registered")
	ResponseWriter(true, "New user registered", nil, http.StatusOK, w)
}

func (hawk *App) login (w http.ResponseWriter, r *http.Request) {
	//get username and password from user
	//check if username exists
	//compare password
	//if matched, set session and current user
	//else return not registered
	formData := User{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		fmt.Println("Could not decode login information")
		ResponseWriter(false, "Could not decode login information", nil, http.StatusInternalServerError, w)
		return
	}
	user := User{}
	//check if username exists
	err = hawk.DB.Where("name = ?", strings.TrimSpace(formData.Username)).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		fmt.Println("User not registered")
		ResponseWriter(false, "User not registered", nil, http.StatusOK, w)
		return
	} else if err != nil {
		fmt.Println("Database error")
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte (strings.TrimSpace(user.Password)), []byte (formData.Password))
	if err != nil {
		fmt.Println("Cannot log in, incorrect password")
		ResponseWriter(false, "Incorrect password, cannot log in", nil, http.StatusUnauthorized, w)
		return
	}
	//username and password have matched
	//setting current user
	currUser := CurrUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Username,
		Access:   user.Access,
		Level:    user.Level,
		Points:   user.Points,
	}
	//set session for 1 day
	err = SetSession(w, currUser, 86400)
	if err != nil {
		ResponseWriter(false, "Error in setting session, user not logged in", nil, http.StatusInternalServerError, w)
		return
	}
	fmt.Println ("User logged in and session set")
	ResponseWriter(true, "User logged in and session is set", nil, http.StatusOK, w)
}

func (hawk *App) logout (w http.ResponseWriter, r *http.Request){
	ClearSession (w)
	fmt.Println("Logout successful")
	ResponseWriter(true, "User logged out", nil, http.StatusOK, w)

}



