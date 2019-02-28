package app

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func ResponseWriter(flag bool, msg string, data interface{}, status int, w http.ResponseWriter) {
	response := Response{
		Success: flag,
		Message: msg,
		Data:    data,
	}

	payload, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)

}

func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	s := "1234567890qwertyuiopasdfghjklzxcvbnmMNBVCXZLKJHFDSAPOIUYTREWQ"
	l := len(s)
	var j int
	s1 := ""
	for i := 0; i < l; i++ {
		j = rand.Intn(l)
		s1 += string(s[j])
	}
	return (s1)

}
//check answer status according to levenstein score
func CheckAnswerStatus(userAnswer string, answer string) int {
	if answer == userAnswer{
		return 1
	} else {
		return 0
	}
}

// GetEnv gets an environment variable with default value as fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
