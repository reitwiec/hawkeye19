package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
	"time"
)

func (hawk *App) addUser(w http.ResponseWriter, r *http.Request) {
	/*
		re := recaptcha.R{
			Secret: "6LftZZoUAAAAAPXZ3nAqHd4jzIbHBNxfMFpuWfMe",
		}
		isValid := re.Verify(*r)
		if isValid {
			fmt.Fprintf(w, "Valid")
		} else {
			//fmt.Fprintf(w, "Invalid! These errors ocurred: %v", re.LastError())
			ResponseWriter(false, "Captcha error", nil, http.StatusBadRequest, w)
			return
		}
	*/
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseWriter(false, "Bad Request", nil, http.StatusBadRequest, w)
		return
	}
	err = validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
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
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Error in hash and salt", nil, http.StatusInternalServerError, w)
		return
	}

	newUser := User{
		Username:       strings.TrimSpace(user.Username),
		Name:           strings.TrimSpace(user.Name),
		Password:       string(hash),
		Access:         0,
		Email:          Sanitize(user.Email),
		Tel:            Sanitize(user.Tel),
		College:        strings.TrimSpace(user.College),
		Region0:        1,
		Region1:        1,
		Region2:        0,
		Region3:        0,
		Region4:        0,
		Region5:        0,
		Banned:         0,
		Points:         0,
		SidequestOrder: SideQuestOrder(),
		UnlockOrder:    UnlockOrder(),
		Timestamp:      time.Now(),
		Country:        strings.TrimSpace(user.Country),
		IsVerified:     0,
		IsMahe:         user.IsMahe,
	}
	//load newUser to database
	tx := hawk.DB.Begin()
	err = tx.Create(&newUser).Error
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok {
			//duplicate entry
			if mysqlErr.Number == 1062 {
				ResponseWriter(false, "Duplicate Entry", nil, http.StatusBadRequest, w)
				return
			}
		}
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error, new user not created", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	token := RandomString()
	//@TODO: send verification email and remove print token statement
	//fmt.Println(token)
	err = SendVUEmail(newUser.Email, token, newUser.Name)
	if err != nil {
		tx.Rollback()
		ResponseWriter(false, "Error in sending verification email", nil, http.StatusInternalServerError, w)
		LogRequest(r, ERROR, err.Error())
		return
	}
	verification := Verification{
		Email: user.Email,
		Token: token,
	}
	//create entry for verification
	err = tx.Create(&verification).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
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
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not decode login information", nil, http.StatusInternalServerError, w)
		return
	}
	user := User{}
	//check if username exists
	err = hawk.DB.Where("username = ?", strings.TrimSpace(formData.Username)).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		ResponseWriter(false, "User not registered", nil, http.StatusOK, w)
		return
	} else if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(strings.TrimSpace(formData.Password)))
	if err != nil {
		LogRequest(r, ERROR, err.Error())
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
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Error in setting session, user not logged in", nil, http.StatusInternalServerError, w)
		return
	}
	user.Password = ""
	ResponseWriter(true, "User logged in and session is set", user, http.StatusOK, w)
}

func (hawk *App) logout(w http.ResponseWriter, r *http.Request) {
	//@TODO: what if user is not logged in ?
	ClearSession(w)
	ResponseWriter(true, "User logged out", nil, http.StatusOK, w)

}

func (hawk *App) forgotPassword(w http.ResponseWriter, r *http.Request) {
	//get email from user
	formData := ForgotPassReq{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		ResponseWriter(false, "Error in decoding form data, password not reset", nil, http.StatusBadRequest, w)
		return
	}
	formData.Email = Sanitize(formData.Email)

	err = validate.Struct(formData)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
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

	user := User{}
	//check if email exists in database
	err = hawk.DB.Where("email = ?", formData.Email).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "User does not exist", nil, http.StatusOK, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//id in forgot_pass_reqs = id in users
	formData.ID = user.ID
	formData.Timestamp = time.Now()
	//generate token for password reset
	token := RandomString()
	//@TODO: Email token and delete print statement
	fmt.Println(token)
	err = SendFPEmail(user.Email, token, user.Name)
	if err != nil {
		ResponseWriter(false, "Error in sending forgot password email", nil, http.StatusInternalServerError, w)
		LogRequest(r, ERROR, err.Error())
		return
	}
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), 14)
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not hash token", nil, http.StatusInternalServerError, w)
		return
	}
	//add to database //update if email already in database
	formData.Token = string(hashedToken)
	tx := hawk.DB.Begin()
	//err = tx.Create(&formData).Error
	//err = tx.Model (&ForgotPassReq{}).Where("email = ?", formData.Email).Update("token, timestamp", formData.Token, formData.Timestamp).Error
	err = tx.Save(&formData).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	/*
		err = SendEmail (formData.Email, formData.Token)
		if err != nil {
			fmt.Println ("Could not email token" + err.Error())
			ResponseWriter(false, "Could not email token", nil, http.StatusInternalServerError, w)
			return
		}
	*/
	ResponseWriter(true, "Password reset link sent", nil, http.StatusOK, w)
	return
}

func (hawk *App) resetPassword(w http.ResponseWriter, r *http.Request) {
	//obtain email from request
	//check if it exists in database
	//compare hash of token and database token
	//if it is same ask for new password and update in db, and delete from forgot_pass_reqs table
	//else abort
	formData := ResetPassReq{}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		ResponseWriter(false, "Could not decode form data", nil, http.StatusBadRequest, w)
		return
	}
	//trim spaces
	formData.Token = strings.TrimSpace(formData.Token)
	formData.Password = strings.TrimSpace(formData.Password)
	formData.Email = strings.TrimSpace(formData.Email)

	forgotPassReqUser := ForgotPassReq{}
	//check if email token exists in database and obtain the associated token
	err = hawk.DB.Where("email = ?", formData.Email).First(&forgotPassReqUser).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Incorrect email", nil, http.StatusOK, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//Email exists in database
	//match tokens
	err = bcrypt.CompareHashAndPassword([]byte(forgotPassReqUser.Token), []byte(formData.Token))
	if err != nil {
		ResponseWriter(false, "Incorrect token", nil, http.StatusOK, w)
		return
	}
	//tokens match
	//if difference in time > 24 hours delete token and return
	//@TODO test this feature
	t24, _ := time.ParseDuration("24h")
	if time.Since(forgotPassReqUser.Timestamp) >= t24 {
		//delete  expired token
		tx := hawk.DB.Begin()
		err = tx.Delete(forgotPassReqUser).Error
		if err != nil {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Expired token, Could not delete expired token", nil, http.StatusInternalServerError, w)
			tx.Rollback()
			return
		}
		tx.Commit()
		ResponseWriter(true, "Expired token, deleted expired token", nil, http.StatusOK, w)
		return
	}
	//create hash for new password
	hash, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	//update the users database
	tx := hawk.DB.Begin()
	err = tx.Model(&User{}).Where("email = ?", forgotPassReqUser.Email).Update("Password", string(hash)).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not update password", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	//delete token entry from forgot_pass_reqs
	err = tx.Delete(&forgotPassReqUser).Error
	//err = tx.Where("email = ?", forgotPassReqUser.Email).Delete(ForgotPassReq{}).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not delete token entry", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	ResponseWriter(true, "Password updated successfully", nil, http.StatusOK, w)
	return
}

func (hawk *App) checkUsername(w http.ResponseWriter, r *http.Request) {
	checkUsername := CheckUsername{}
	err := json.NewDecoder(r.Body).Decode(&checkUsername)
	if err != nil {
		ResponseWriter(false, "Error in decoding", nil, http.StatusBadRequest, w)
		return
	}
	//look for username in string
	checkUsername.Username = Sanitize(checkUsername.Username)
	err = hawk.DB.Where("Username = ?", checkUsername.Username).Find(&User{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(true, "Username available", nil, http.StatusOK, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//username exists
	ResponseWriter(false, "Username taken", nil, http.StatusOK, w)
	return
}

func (hawk *App) checkEmail(w http.ResponseWriter, r *http.Request) {
	checkEmail := CheckEmail{}
	err := json.NewDecoder(r.Body).Decode(&checkEmail)
	if err != nil {
		ResponseWriter(false, "Error in decoding", nil, http.StatusBadRequest, w)
		return
	}
	//look for email in DB
	checkEmail.Email = Sanitize(checkEmail.Email)
	err = hawk.DB.Where("Email = ?", checkEmail.Email).Select("email").Find(&User{}).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(true, "Email available", nil, http.StatusOK, w)
			return
		} else {
			LogRequest(r, ERROR, err.Error())
			ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
			return
		}
	}
	//username exists
	ResponseWriter(false, "Email registered", nil, http.StatusOK, w)
	return
}

func (hawk *App) verifyUser(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["email"]
	if !ok || len(keys) < 1 {
		ResponseWriter(false, "Bad request", nil, http.StatusBadRequest, w)
		return
	}
	email := keys[0]
	keys, ok = r.URL.Query()["token"]
	if !ok || len(keys) < 1 {
		ResponseWriter(false, "Bad request", nil, http.StatusBadRequest, w)
		return
	}
	token := keys[0]
	verificationData := Verification{}
	err := hawk.DB.Where("email = ? AND token = ? ", email, token).First(&verificationData).Error
	/*
		if err == nil {
			fmt.Println(verificationData)
		}
	*/
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Wrong data", nil, http.StatusOK, w)
			return
		}
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	//token exists
	tx := hawk.DB.Begin()
	err = tx.Where("email = ?", email).First(&User{}).Update("is_verified", 1).Error
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		tx.Rollback()
		return
	}
	tx.Commit()
	ResponseWriter(true, "User verified", nil, http.StatusOK, w)
}
