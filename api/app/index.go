package app

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"os"
)

var (
	Configuration Settings
	validate      *validator.Validate
	CookieHandler *securecookie.SecureCookie
)

func (hawk *App) Initialise() {
	//initialise logger
	//create your file with desired read/write permissions
	var err error
	hawk.logFile, err = os.OpenFile("info.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//set output of logs to f
	log.SetOutput(hawk.logFile)
	//test case
	log.Println("Logging to file")
	//configure Instance
	ConfigureInstance()
	log.Printf(`

Configuration:
DB Username: %s
DB Password: %s
DB Name: %s
Hash Key: %s
Block Key: %s

`, Configuration.DBUsername, Configuration.DBPassword, Configuration.DBName, Configuration.HashKey, Configuration.BlockKey)

	//configure mux
	hawk.LoadRoutes()
	//initialise validator
	validate = validator.New()
	//DB connection string
	Configuration.DBConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		Configuration.DBUsername, Configuration.DBPassword, Configuration.DBName)

	//create securecookie instance
	CookieHandler = securecookie.New([]byte(Configuration.HashKey), []byte(Configuration.BlockKey))

}

func (hawk *App) migrate() {
	hawk.DB.CreateTable(
		&Help{},
		&User{},
		&Attempt{},
		&Question{},
		&Hint{},
		&ForgotPassReq{},
		&Verification{},
	)
}

func (hawk *App) Run(Args []string) {
	var err error
	hawk.DB, err = gorm.Open("mysql", Configuration.DBConn)
	if err != nil {
		log.Fatalf("Could not establish database connection, shutting down\n\t" + err.Error())
		return
	}
	log.Println("Database connection established")
	defer hawk.DB.Close()

	for _, arg := range Args {
		switch arg {
		case "-m":
			fmt.Println("migrate entered")
			hawk.migrate()
		}
	}

	err = http.ListenAndServe(Configuration.ServerAddr, hawk.router)
	if err != nil {
		log.Fatalf("Could not start server, shutting down")
	}

	defer hawk.logFile.Close()

}
