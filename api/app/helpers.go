package app

import (
	"encoding/json"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	SimilarStringRatio = 0.75
	CorrectAnswer      = 1
	CloseAnswer        = 2
	IncorrectAnswer    = 3
	LevelCount		   = 6
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
	ratio := levenshtein.RatioForStrings([]rune(userAnswer), []rune(answer), levenshtein.DefaultOptions)
	if ratio == 1 {
		return CorrectAnswer
	}
	if ratio >= 0.75 {
		return CloseAnswer
	}
	return IncorrectAnswer
}

// GetEnv gets an environment variable with default value as fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Sanitize(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	return s
}
/*
func SideQuestOrdere () [LevelCount] int {
	var index int
	a := [LevelCount]int{1,2,3,4,5,7}
	//shuffle a
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i:= 0 ; i < LevelCount; i++ {
		//generate random index
		index = r1.Intn(LevelCount)
		//swap ith number with number at index
		a[i], a[index] = a[index], a[i]
	}
	return a
}
*/
func SideQuestOrder () string {
	permutations := [120] string {
		"1,2,3,4,5,6",
		"1,3,2,4,5,6",
		"1,4,2,3,5,6",
		"1,2,4,3,5,6",
		"1,3,4,2,5,6",
		"1,4,3,2,5,6",
		"1,4,3,5,2,6",
		"1,3,4,5,2,6",
		"1,5,4,3,2,6",
		"1,4,5,3,2,6",
		"1,3,5,4,2,6",
		"1,5,3,4,2,6",
		"1,5,2,4,3,6",
		"1,2,5,4,3,6",
		"1,4,5,2,3,6",
		"1,5,4,2,3,6",
		"1,2,4,5,3,6",
		"1,4,2,5,3,6",
		"1,3,2,5,4,6",
		"1,2,3,5,4,6",
		"1,5,3,2,4,6",
		"1,3,5,2,4,6",
		"1,2,5,3,4,6",
		"1,5,2,3,4,6",
		"1,6,2,3,4,5",
		"1,2,6,3,4,5",
		"1,3,6,2,4,5",
		"1,6,3,2,4,5",
		"1,2,3,6,4,5",
		"1,3,2,6,4,5",
		"1,3,2,4,6,5",
		"1,2,3,4,6,5",
		"1,4,3,2,6,5",
		"1,3,4,2,6,5",
		"1,2,4,3,6,5",
		"1,4,2,3,6,5",
		"1,4,6,3,2,5",
		"1,6,4,3,2,5",
		"1,3,4,6,2,5",
		"1,4,3,6,2,5",
		"1,6,3,4,2,5",
		"1,3,6,4,2,5",
		"1,2,6,4,3,5",
		"1,6,2,4,3,5",
		"1,4,2,6,3,5",
		"1,2,4,6,3,5",
		"1,4,6,2,3,5",
		"1,6,4,2,3,5",
		"1,5,6,2,3,4",
		"1,6,5,2,3,4",
		"1,2,5,6,3,4",
		"1,5,2,6,3,4",
		"1,6,2,5,3,4",
		"1,2,6,5,3,4",
		"1,2,6,3,5,4",
		"1,6,2,3,5,4",
		"1,3,2,6,5,4",
		"1,2,3,6,5,4",
		"1,6,3,2,5,4",
		"1,3,6,2,5,4",
		"1,3,5,2,6,4",
		"1,5,3,2,6,4",
		"1,2,3,5,6,4",
		"1,3,2,5,6,4",
		"1,5,2,3,6,4",
		"1,2,5,3,6,4",
		"1,6,5,3,2,4",
		"1,5,6,3,2,4",
		"1,3,6,5,2,4",
		"1,6,3,5,2,4",
		"1,5,3,6,2,4",
		"1,3,5,6,2,4",
		"1,4,5,6,2,3",
		"1,5,4,6,2,3",
		"1,6,4,5,2,3",
		"1,4,6,5,2,3",
		"1,5,6,4,2,3",
		"1,6,5,4,2,3",
		"1,6,5,2,4,3",
		"1,5,6,2,4,3",
		"1,2,6,5,4,3",
		"1,6,2,5,4,3",
		"1,5,2,6,4,3",
		"1,2,5,6,4,3",
		"1,2,4,6,5,3",
		"1,4,2,6,5,3",
		"1,6,2,4,5,3",
		"1,2,6,4,5,3",
		"1,4,6,2,5,3",
		"1,6,4,2,5,3",
		"1,5,4,2,6,3",
		"1,4,5,2,6,3",
		"1,2,5,4,6,3",
		"1,5,2,4,6,3",
		"1,4,2,5,6,3",
		"1,2,4,5,6,3",
		"1,3,4,5,6,2",
		"1,4,3,5,6,2",
		"1,5,3,4,6,2",
		"1,3,5,4,6,2",
		"1,4,5,3,6,2",
		"1,5,4,3,6,2",
		"1,5,4,6,3,2",
		"1,4,5,6,3,2",
		"1,6,5,4,3,2",
		"1,5,6,4,3,2",
		"1,4,6,5,3,2",
		"1,6,4,5,3,2",
		"1,6,3,5,4,2",
		"1,3,6,5,4,2",
		"1,5,6,3,4,2",
		"1,6,5,3,4,2",
		"1,3,5,6,4,2",
		"1,5,3,6,4,2",
		"1,3,4,6,5,2",
		"1,4,3,6,5,2",
		"1,6,4,3,5,2",
		"1,4,6,3,5,2",
		"1,3,6,4,5,2",
		"1,6,3,4,5,2",
	}
	rand.Seed(time.Now().UnixNano())
	//random integer to pick random string
	n := rand.Intn(120)
	return permutations[n]

}