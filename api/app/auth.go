package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func (hawk *App) addUser(w http.ResponseWriter, r *http.Request) {
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
		var fields []string
		for _, err := range err.(validator.ValidationErrors) {
			fields = append(fields, err.Field())
		}
		ResponseWriter(false, "Bad Request", fields, http.StatusBadRequest, w)
		return
	}

	//validation successful
	//hash  and salt password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println(err)
		ResponseWriter(false, "Error in hash and salt", nil, http.StatusInternalServerError, w)
		return
	}

	unlockOrder := ""
	count := 1

	perm := rand.Perm(3)
	fmt.Println(perm)
	for i := range perm {
		fmt.Println(i)
		i = perm[i] + 2
		unlockOrder += strconv.Itoa(i)
		if count != len(perm) {
			unlockOrder += ":"
		}
		count++
	}

	newUser := User{
		Username:    Sanitize(user.Username),
		Name:        Sanitize(user.Name),
		Password:    string(hash),
		Access:      0,
		Email:       Sanitize(user.Email),
		Tel:         Sanitize(user.Tel),
		College:     Sanitize(user.College),
		Region1:     1,
		Region2:     0,
		Region3:     0,
		Region4:     0,
		Region5:     0,
		Banned:      0,
		Points:      0,
		SideQuest:   SideQuestOrder(),
		UnlockOrder: unlockOrder,
	}
	//load newUser to database
	//begin transaction
	tx := hawk.DB.Begin()
	err = tx.Create(&newUser).Error
	if err != nil {
		fmt.Println("Database error, new user not created")
		ResponseWriter(false, "Database error, new user not created", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	fmt.Println("New user registered")
	ResponseWriter(true, "New user registered", nil, http.StatusOK, w)
}

func (hawk *App) login(w http.ResponseWriter, r *http.Request) {
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
	// @TODO: add recordnotfound to other DB queries as well
	err = hawk.DB.Where("username = ?", Sanitize(formData.Username)).First(&user).Error
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Sanitize(formData.Password)))
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
	}
	//set session for 1 day
	err = SetSession(w, currUser, 86400)
	if err != nil {
		ResponseWriter(false, "Error in setting session, user not logged in", nil, http.StatusInternalServerError, w)
		return
	}
	fmt.Println("User logged in and session set")
	ResponseWriter(true, "User logged in and session is set", currUser, http.StatusOK, w)
}

func (hawk *App) logout(w http.ResponseWriter, r *http.Request) {
	//@TODO: what if user is not logged in ?
	ClearSession(w)
	fmt.Println("Logout successful")
	ResponseWriter(true, "User logged out", nil, http.StatusOK, w)

}

func (hawk *App) forgotPassword(w http.ResponseWriter, r *http.Request) {
	//get email from user
	formData := ForgotPassReq{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		fmt.Println("Error in decoding form data")
		ResponseWriter(false, "Error in decoding form data, password not reset", nil, http.StatusBadRequest, w)
		return
	}
	formData.Email = Sanitize(formData.Email)
	err = validate.Struct(formData)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			ResponseWriter(false, "Invalid validation error ", nil, http.StatusBadRequest, w)
			return
		}
		var fields []string
		for _, err := range err.(validator.ValidationErrors) {
			fields = append(fields, err.Field())
		}
		ResponseWriter(false, "Bad Request", fields, http.StatusBadRequest, w)
		return
	}

	//check if email exists in database
	err = hawk.DB.Where("email = ?", formData.Email).First(&formData).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Email not found")
			ResponseWriter(false, "User does not exist", nil, http.StatusOK, w)
			return
		} else {
			fmt.Println("Database error")
			fmt.Println(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}

	formData.Timestamp = time.Now()
	//generate token for password reset
	token := RandomString()
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), 14)
	if err != nil {
		fmt.Println("Could not hash token")
		ResponseWriter(false, "Could not hash token", nil, http.StatusInternalServerError, w)
		return
	}
	//add to database
	formData.ID = string(hashedToken)
	tx := hawk.DB.Begin()
	err = tx.Create(&formData).Error
	if err != nil {
		fmt.Println("Database error")
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	//@TODO: Email the token
	fmt.Println(token)
	ResponseWriter(true, "Password reset link sent", nil, http.StatusOK, w)
	return
}

func (hawk *App) resetPassword(w http.ResponseWriter, r *http.Request) {
	//obtain id from request
	//hash it
	//check if it exists in database
	//if it does ask for new password and update in db, and delete from password reset table
	//else abort
	//obtain token from request
	formData := ResetPassReq{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		fmt.Println("Could not decode json data")
		ResponseWriter(false, "Could not decode form data", nil, http.StatusInternalServerError, w)
		return
	}
	//trim spaces
	formData.Token = Sanitize(formData.Token)
	formData.Password = Sanitize(formData.Password)
	//hash the token
	hash, err := bcrypt.GenerateFromPassword([]byte(formData.Token), 14)
	hashedToken := string(hash)
	forgotPassReqUser := ForgotPassReq{}
	//check if hashed token exists in database
	err = hawk.DB.Where("ID = ?", hashedToken).First(&forgotPassReqUser).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Incorrect token")
			ResponseWriter(false, "Incorrect token", nil, http.StatusOK, w)
			return
		} else {
			fmt.Println("Database error")
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//hashed token exists in database
	//if difference in time > 24 hours delete token and return
	t24, _ := time.ParseDuration("24h")
	if time.Since(forgotPassReqUser.Timestamp) >= t24 {
		fmt.Println("Token expired")
		ResponseWriter(false, "Token Expired", nil, http.StatusOK, w)
		//delete token
		tx := hawk.DB.Begin()
		err = tx.Delete(forgotPassReqUser).Error
		if err != nil {
			fmt.Println("Could not delete expired token")
			ResponseWriter(false, "Could not delete expired token", nil, http.StatusOK, w)
			tx.Rollback()
			return
		}
		tx.Commit()
		ResponseWriter(true, "Deleted expired token", nil, http.StatusOK, w)
		return
	}
	//create hash for new password
	hash, err = bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	//update the users database
	forgotPassUser := User{}
	hawk.DB.Where("Email = ? ", forgotPassReqUser.Email).First(&forgotPassUser)
	tx := hawk.DB.Begin()
	err = tx.Model(&forgotPassUser).Update("Password", string(hash)).Error
	if err != nil {
		fmt.Println("Could not update password")
		ResponseWriter(false, "Could not update password", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	tx.Begin()
	//delete token entry from forgotpassrequser
	err = tx.Delete(&forgotPassReqUser).Error
	if err != nil {
		fmt.Println("Could not delete token entry but password updated")
		ResponseWriter(false, "Could not delete token entry but password updated", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	fmt.Println("Password updated successfully")
	ResponseWriter(true, "Password updated successfully", nil, http.StatusOK, w)
	return
}

func (hawk *App) checkUsername(w http.ResponseWriter, r *http.Request) {
	checkUsername := CheckUsername{}
	err := json.NewDecoder(r.Body).Decode(&checkUsername)
	if err != nil {
		fmt.Println("Error in decoding")
		ResponseWriter(false, "Error in decoding", nil, http.StatusInternalServerError, w)
		return
	}
	//look for username in string
	checkUsername.Username = Sanitize(checkUsername.Username)
	err = hawk.DB.Where("Username = ?", checkUsername.Username).Find(&User{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Username available")
			ResponseWriter(true, "Username available", nil, http.StatusOK, w)
			return
		} else {
			fmt.Println("Database error")
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//username exists
	fmt.Println("Username taken")
	ResponseWriter(false, "Username taken", nil, http.StatusOK, w)
	return
}

func (hawk *App) checkEmail(w http.ResponseWriter, r *http.Request) {
	checkEmail := CheckEmail{}
	err := json.NewDecoder(r.Body).Decode(&checkEmail)
	if err != nil {
		fmt.Println("Error in decoding")
		ResponseWriter(false, "Error in decoding", nil, http.StatusInternalServerError, w)
		return
	}
	//look for email in DB
	checkEmail.Email = Sanitize(checkEmail.Email)
	err = hawk.DB.Where("Email = ?", checkEmail.Email).Select("email").Find(&User{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Email available")
			ResponseWriter(true, "Email available", nil, http.StatusOK, w)
			return
		} else {
			fmt.Println("Database error")
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//username exists
	fmt.Println("Email registered")
	ResponseWriter(false, "Email registered", nil, http.StatusOK, w)
	return
}
