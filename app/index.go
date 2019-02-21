package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
)

var (
	Configuration	Settings
	validate		*validator.Validate
)

func (hawk *App) Initialise () {
	//configure Instance
	ConfigureInstance ()
	//configure mux
	hawk.LoadRoutes ()
	//initialise validator
	validate = validator.New ()
	//DB connection string
	Configuration.DBConn = fmt.Sprintf ("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		Configuration.DBUsername, Configuration.DBPassword, Configuration.DBName)



}

func (hawk *App) Run (Args [] string) {
	var err error
	hawk.DB, err = gorm.Open ("mysql", Configuration.DBConn)
	if err != nil {
		log.Fatalf ("Could not establish database connection, shutting down")
		return
	}
	fmt.Println ("Database connection established")
	defer hawk.DB.Close ()
	fmt.Println ("Server started")
	err = http.ListenAndServe(Configuration.ServerAddr, hawk.router)
	if err != nil {
		fmt.Println ("Could not start server, shutting down")
	}


}