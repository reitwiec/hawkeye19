package app

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New (Configuration.HashKey, Configuration.BlockKey)

func SetSession (w http.ResponseWriter, currUser CurrUser, age int) error {

	if encoded, err := cookieHandler.Encode ("session", currUser); err == nil {
		cookie := &http.Cookie {
			Name: "session",
			Value: encoded,
			Path: "/",
			MaxAge: age,
		}
		http.SetCookie (w, cookie)
	} else {
		fmt.Println ("Some error in setting session\n\t" + err.Error ())
		return err
	}
	fmt.Println("Cookie set!")
	return nil
}


func ClearSession (w http.ResponseWriter) {
	cookie := &http.Cookie {
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}