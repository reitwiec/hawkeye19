package app

import (
	"encoding/json"
	"fmt"
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
	LevelCount         = 6
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
func SideQuestOrder() string {
	permutations := [120]string{
		"123456",
		"132456",
		"142356",
		"124356",
		"134256",
		"143256",
		"143526",
		"134526",
		"154326",
		"145326",
		"135426",
		"153426",
		"152436",
		"125436",
		"145236",
		"154236",
		"124536",
		"142536",
		"132546",
		"123546",
		"153246",
		"135246",
		"125346",
		"152346",
		"162345",
		"126345",
		"136245",
		"163245",
		"123645",
		"132645",
		"132465",
		"123465",
		"143265",
		"134265",
		"124365",
		"142365",
		"146325",
		"164325",
		"134625",
		"143625",
		"163425",
		"136425",
		"126435",
		"162435",
		"142635",
		"124635",
		"146235",
		"164235",
		"156234",
		"165234",
		"125634",
		"152634",
		"162534",
		"126534",
		"126354",
		"162354",
		"132654",
		"123654",
		"163254",
		"136254",
		"135264",
		"153264",
		"123564",
		"132564",
		"152364",
		"125364",
		"165324",
		"156324",
		"136524",
		"163524",
		"153624",
		"135624",
		"145623",
		"154623",
		"164523",
		"146523",
		"156423",
		"165423",
		"165243",
		"156243",
		"126543",
		"162543",
		"152643",
		"125643",
		"124653",
		"142653",
		"162453",
		"126453",
		"146253",
		"164253",
		"154263",
		"145263",
		"125463",
		"152463",
		"142563",
		"124563",
		"134562",
		"143562",
		"153462",
		"135462",
		"145362",
		"154362",
		"154632",
		"145632",
		"165432",
		"156432",
		"146532",
		"164532",
		"163542",
		"136542",
		"156342",
		"165342",
		"135642",
		"153642",
		"134652",
		"143652",
		"164352",
		"146352",
		"136452",
		"163452",
	}
	rand.Seed(time.Now().UnixNano())
	//random integer to pick random string
	n := rand.Intn(120)
	return permutations[n]
}

func UnlockOrder() string {
	permutations := [6]string{
		"0,2,3,4,5",
		"0,2,4,3,5",
		"0,3,2,4,5",
		"0,3,4,2,5",
		"0,4,2,3,5",
		"0,4,3,2,5",
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6)
	return permutations[n]
}

func GetNextRegion(unlockOrder string) uint8 {
	//traverse till non 0 number is encountered
	l := len(unlockOrder)
	for i := 0; i < l; i++ {
		if unlockOrder[i] != ',' && unlockOrder[i] != '0' {
			fmt.Println(unlockOrder[i])
			return unlockOrder[i]
		}
	}
	return '0' //error
}

func UpdateUnlockOrder(unlockOrder string, region int) string {
	l := len(unlockOrder)
	var updatedUO string
	aregion := uint8(region + 48)
	for i := 0; i < l; i++ {
		if unlockOrder[i] == aregion {
			updatedUO = unlockOrder[:i] + "0" + unlockOrder[i+1:]
			break
		}
	}
	return updatedUO
}
