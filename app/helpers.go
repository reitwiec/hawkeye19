package app

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func ResponseWriter (flag bool, msg string, data interface {}, status int, w http.ResponseWriter){
	response := Response{
		Success: flag,
		Message: msg,
		Data: data,
	}

	payload, err := json.Marshal (response)

	if err !=  nil {
		http.Error (w, err.Error (), http.StatusInternalServerError)
		return
	}

	w.Header ().Set ("Content-Type", "application/json")
	w.Write (payload)

}

func convertStringToByteSlice(s string) []byte {
	x := strings.Split(s, ",")
	b := make([]byte, len(x))
	for i := range x {
		y, _ := strconv.Atoi(x[i])
		b[i] = byte(y)
	}
	return b
}

func RandomString () string {
	s := "1234567890qwertyuiopasdfghjklzxcvbnmMNBVCXZLKJHFDSAPOIUYTREWQ"
	l := len (s)
	var j int
	s1 := ""
	for i := 0; i<l; i++ {
		j = rand.Intn(l)
		s1 +=  string (s[j])
	}
	return (s1)

}