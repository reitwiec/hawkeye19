package app

import (
	"fmt"
	"net/http"
)

func SetSession(w http.ResponseWriter, currUser CurrUser, age int) error {

	if encoded, err := CookieHandler.Encode("session", currUser); err == nil {
		cookie := &http.Cookie{
			Name:   "session",
			Value:  encoded,
			Path:   "/",
			MaxAge: age,
		}
		http.SetCookie(w, cookie)
	} else {
		fmt.Println("Some error in setting session\n\t" + err.Error())
		return err
	}
	fmt.Println("Cookie set!")
	return nil
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func GetCurrUser(w http.ResponseWriter, r *http.Request) (CurrUser, error) {
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Error in reading cookie\n\t" + err.Error())
		return CurrUser{}, err
	}
	value := CurrUser{}
	if err = CookieHandler.Decode("session", c.Value, &value); err == nil {
		//ResponseWriter(true, "Returning current user", value, http.StatusOK, w)
		fmt.Println("Returning curruser")
		return value, nil
	} else {
		//ResponseWriter(false, "Could not read cookie data", nil, http.StatusInternalServerError, w)
		fmt.Println("Could not read cookie data\n" + err.Error())
		return CurrUser{}, err
	}
}
