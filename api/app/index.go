package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/go-playground/validator.v9"
)

var (
	Configuration Settings
	validate      *validator.Validate
	CookieHandler *securecookie.SecureCookie
)

func (hawk *App) Initialise() {
	//configure Instance
	ConfigureInstance()
	fmt.Printf(`

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
	fmt.Println("Creating Tables")
	hawk.DB.CreateTable(
		&User{},
		&Attempt{},
		&Question{},
		&Hint{},
		&ForgotPassReq{},
	)
}

func (hawk *App) Run(Args []string) {
	var err error
	hawk.DB, err = gorm.Open("mysql", Configuration.DBConn)
	if err != nil {
		log.Fatalf("Could not establish database connection, shutting down\n\t" + err.Error())
		return
	}
	fmt.Println("Database connection established")
	defer hawk.DB.Close()

	for _, arg := range Args {
		switch arg {
		case "-m":
			hawk.migrate()
		}
	}

	fmt.Println("Server started")
	err = http.ListenAndServe(Configuration.ServerAddr, hawk.router)
	if err != nil {
		fmt.Println("Could not start server, shutting down")
	}

}
