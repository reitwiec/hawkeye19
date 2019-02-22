package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"time"
)

type Settings struct {
	ServerAddr	string	`json:"ServerAddr"`

	DBUsername	string	`json:"DBUsername"`
	DBPassword	string	`json:"DBPassword"`
	DBName		string 	`json:"DBName"`
	DBConn		string
	HashKey		string	`json:"HashKey"`
	BlockKey	string	`json:"BlockKey"`
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
	Level    int
	Points	 int
}

type App struct {
	DB       *gorm.DB
	router   *mux.Router
	currUser CurrUser
}

type User struct {
	ID			int		`gorm:"primary_key;auto_increment" json:"userID"`
	Name 		string 	`gorm:"not null" json:"name" validate:"required"`
	Username	string	`gorm:"not null;unique" json:"username" validate:"alphanum,required"`
	Password	string	`gorm:"not null" json:"password" validate:"min=8,max=24,required"`
	Access		int	`json:"access"`

	//general info
	Email		string	`gorm:"unique;not null" json:"email" validate:"email,required"`
	Tel			string	`gorm:"not null;unique" json:"tel" validate:"len=10,required"`
	College		string	`gorm:"not null" json:"college" validate:"alpha,required"`

	//Gameplay status
	Level 		int		`json:"level"`
	Banned		int 	`json:"banned"`
	Points 		int		`json:"points"`
}

type ForgotPassReq struct {
	ID		int	`gorm:"auto_increment;primary_key" json:"id;omitempty"`
	Email	string	`gorm:"not null" json:"email" validate:"email"`
	Token 	string	`json:"token"`
	Timestamp 	time.Time	`gorm:"" json:"timestamp"`
}

type ResetPassReq struct {
	Token	string	`json:"token" validate:"required"`
	Password	string	`json:"password" validate:"min=8,max=24,required"`
}