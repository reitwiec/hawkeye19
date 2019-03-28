package app

import (
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Settings struct {
	ServerAddr string `json:"ServerAddr"`

	DBUsername   string `json:"DBUsername"`
	DBPassword   string `json:"DBPassword"`
	DBName       string `json:"DBName"`
	DBConn       string
	HashKey      string `json:"HashKey"`
	BlockKey     string `json:"BlockKey"`
	RecaptchaKey string `json:"RecaptchaKey"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type CurrUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type App struct {
	DB       *gorm.DB
	router   *mux.Router
	currUser CurrUser
	logFile  *os.File
}

type User struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"userID"`
	Name      string    `gorm:"not null" json:"name" validate:"required"`
	Username  string    `gorm:"not null;unique" json:"username" validate:"alphanum,required"`
	Password  string    `gorm:"not null" json:"password" validate:"min=8,max=24,required"`
	Access    int       `json:"access"`
	Banned    int       `json:"banned"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`

	//general info
	Email      string `gorm:"unique;not null" json:"email" validate:"email,required"`
	Tel        string `gorm:"not null;unique" json:"tel" validate:"len=10,required"`
	College    string `gorm:"not null" json:"college" validate:"required"`
	Country    string `gorm:"not null" json:"country"`
	IsMahe     int    `gorm:"not null" json:"isMahe"` //1 if MAHE else, 0
	IsVerified int    `gorm:"not null" json:"isVerified"`
	FirstLogin int    `gorm:"not null" json:"firstLogin"` //1 by default, sets to 0 on first login

	//Gameplay status
	Region0 int `json:"region0" ` //sidequest
	Region1 int `json:"region1" `
	Region2 int `json:"region2" `
	Region3 int `json:"region3" `
	Region4 int `json:"region4" `
	Region5 int `json:"region5"` //linear

	Points          int    `json:"points"`
	SidequestOrder  string `json:"sideQuestOrder"`
	UnlockOrder     string `json:"unlockOrder"`
	SideQuestPoints int    `json:"sideQuestPoints"`
}

type Help struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"helpID"`
	User      int       `gorm:"not null" json:"userID" validate:"required"`
	Username  string    `gorm:"not null" json:"username" validate:"required"`
	Message   string    `gorm:"not null" validate:"required"`
	Email     string    `gorm:"not null" validate:"email,required"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}

type Attempt struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"attemptID"`
	User      int       `gorm:"not null" json:"userID" validate:"required"`
	Question  int       `gorm:"not null" json:"questionID" validate:"required"`
	Answer    string    `gorm:"not null" json:"answer" validate:"alphanum; required"`
	Status    int       `gorm:"not null" json:"status" validate:"required"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}

type Question struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"questionID"`
	Level     int       `gorm:"not null" json:"level" validate:"required"`
	Region    int       `gorm:"not null" json:"region" validate:"required"`
	Question  string    `gorm:"not null" json:"question" validate:"required"`
	Answer    string    `gorm:"not null" json:"answer,omitempty" validate:"required"`
	AddInfo   string    `json:"addinfo"`
	AddedBy   string    `gorm:"not null" json:"addedBy" validate:"required"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}

type Hint struct {
	ID        int       `gorm:"primary_key; autoincrement" json:"hintID"`
	Question  int       `gorm:"not null" json:"questionID" validate:"required"`
	Active    int       `gorm:"not null" json:"active"`
	Hint      string    `json:"hint" validate:"required"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
	AddedBy   string    `gorm:"not null" json:"addedBy" validate:"required"`
}

type ForgotPassReq struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Token     string    `gorm:"not null" json:"token"`
	Email     string    `gorm:"not null;unique" json:"email" validate:"email"`
	Timestamp time.Time `json:"timestamp"`
}

type ResetPassReq struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckUsername struct {
	Username string
}

type CheckEmail struct {
	Email string
}

type Verification struct {
	ID    int `gorm:"primary_key;auto_increment" json:"id"`
	Email string
	Token string
}

type LogInfo struct {
	Timestamp time.Time
	Method    string
	URL       string
	Body      string
	User      interface{}
	Response  *Response
}
