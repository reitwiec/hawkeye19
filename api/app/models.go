package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"time"
)

type Settings struct {
	ServerAddr string `json:"ServerAddr"`

	DBUsername string `json:"DBUsername"`
	DBPassword string `json:"DBPassword"`
	DBName     string `json:"DBName"`
	DBConn     string
	HashKey    string `json:"HashKey"`
	BlockKey   string `json:"BlockKey"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type CurrUser struct {
	ID       int
	Username string
	Email    string
	Access   int
	Region1  int
	Region2  int
	Region3  int
	Region4  int
	Region5  int
	Points   int
}

type App struct {
	DB       *gorm.DB
	router   *mux.Router
	currUser CurrUser
}

type User struct {
	ID       int    `gorm:"primary_key;auto_increment" json:"userID"`
	Name     string `gorm:"not null" json:"name" validate:"required"`
	Username string `gorm:"not null;unique" json:"username" validate:"alphanum,required"`
	Password string `gorm:"not null" json:"password" validate:"min=8,max=24,required"`
	Access   int    `json:"access"`

	//general info
	Email   string `gorm:"unique;not null" json:"email" validate:"email,required"`
	Tel     string `gorm:"not null;unique" json:"tel" validate:"len=10,required"`
	College string `gorm:"not null" json:"college" validate:"alpha,required"`

	//Gameplay status
	Region1   int    `json:"Region1"`
	Region2   int    `json:"Region2"`
	Region3   int    `json:"Region3"`
	Region4   int    `json:"Region4"`
	Region5   int    `json:"Region5"`
	Banned    int    `json:"banned"`
	Points    int    `json:"points"`
	Sidequest string `json:"sidequest"`
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
	ID       int    `gorm:"primary_key;auto_increment" json:"questionID"`
	Level    int    `gorm:"not null" json:"level" validate:"required"`
	Region   int    `gorm:"not null" json:"region" validate:"required"`
	Question string `gorm:"not null" json:"question" validate:"required"`
	Answer   string `gorm:"not null" json:"answer,omitempty" validate:"required"`
	AddInfo  string `json:"addinfo"`
	AddedBy  string `gorm:"not null" json:"addedBy" validate:"required"`
}

type Hint struct {
	ID       int    `gorm:"primary_key; autoincrement" json:"hintID"`
	Question int    `gorm:"not null" json:"questionID" validate:"required"`
	Active   int    `gorm:"not null" json:"active"`
	Hint     string `json:"hint" validate:"required"`
}

type ForgotPassReq struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"not null" json:"email" validate:"email"`
	Timestamp time.Time `json:"timestamp"`
}

type ResetPassReq struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type CheckUsername struct {
	Username	string
}

type CheckEmail struct {
	Email	string
}