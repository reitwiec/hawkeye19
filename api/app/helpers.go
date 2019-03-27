package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"io/ioutil"
	"log"
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
	ERROR              = "ERROR"
	INFO               = "INFO"
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
	/*
		//if status == 500 error, log it
		log.Printf("ERROR\nSuccess: %t\nMessage: %s\nData: %s\n\n",
			response.Success, response.Message, response.Data)
	*/

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
		"1,2,3,4,5,6",
		"1,2,3,4,6,5",
		"1,2,3,5,4,6",
		"1,2,3,5,6,4",
		"1,2,3,6,4,5",
		"1,2,3,6,5,4",
		"1,2,4,3,5,6",
		"1,2,4,3,6,5",
		"1,2,4,5,3,6",
		"1,2,4,5,6,3",
		"1,2,4,6,3,5",
		"1,2,4,6,5,3",
		"1,2,5,3,4,6",
		"1,2,5,3,6,4",
		"1,2,5,4,3,6",
		"1,2,5,4,6,3",
		"1,2,5,6,3,4",
		"1,2,5,6,4,3",
		"1,2,6,3,4,5",
		"1,2,6,3,5,4",
		"1,2,6,4,3,5",
		"1,2,6,4,5,3",
		"1,2,6,5,3,4",
		"1,2,6,5,4,3",
		"1,3,2,4,5,6",
		"1,3,2,4,6,5",
		"1,3,2,5,4,6",
		"1,3,2,5,6,4",
		"1,3,2,6,4,5",
		"1,3,2,6,5,4",
		"1,3,4,2,5,6",
		"1,3,4,2,6,5",
		"1,3,4,5,2,6",
		"1,3,4,5,6,2",
		"1,3,4,6,2,5",
		"1,3,4,6,5,2",
		"1,3,5,2,4,6",
		"1,3,5,2,6,4",
		"1,3,5,4,2,6",
		"1,3,5,4,6,2",
		"1,3,5,6,2,4",
		"1,3,5,6,4,2",
		"1,3,6,2,4,5",
		"1,3,6,2,5,4",
		"1,3,6,4,2,5",
		"1,3,6,4,5,2",
		"1,3,6,5,2,4",
		"1,3,6,5,4,2",
		"1,4,2,3,5,6",
		"1,4,2,3,6,5",
		"1,4,2,5,3,6",
		"1,4,2,5,6,3",
		"1,4,2,6,3,5",
		"1,4,2,6,5,3",
		"1,4,3,2,5,6",
		"1,4,3,2,6,5",
		"1,4,3,5,2,6",
		"1,4,3,5,6,2",
		"1,4,3,6,2,5",
		"1,4,3,6,5,2",
		"1,4,5,2,3,6",
		"1,4,5,2,6,3",
		"1,4,5,3,2,6",
		"1,4,5,3,6,2",
		"1,4,5,6,2,3",
		"1,4,5,6,3,2",
		"1,4,6,2,3,5",
		"1,4,6,2,5,3",
		"1,4,6,3,2,5",
		"1,4,6,3,5,2",
		"1,4,6,5,2,3",
		"1,4,6,5,3,2",
		"1,5,2,3,4,6",
		"1,5,2,3,6,4",
		"1,5,2,4,3,6",
		"1,5,2,4,6,3",
		"1,5,2,6,3,4",
		"1,5,2,6,4,3",
		"1,5,3,2,4,6",
		"1,5,3,2,6,4",
		"1,5,3,4,2,6",
		"1,5,3,4,6,2",
		"1,5,3,6,2,4",
		"1,5,3,6,4,2",
		"1,5,4,2,3,6",
		"1,5,4,2,6,3",
		"1,5,4,3,2,6",
		"1,5,4,3,6,2",
		"1,5,4,6,2,3",
		"1,5,4,6,3,2",
		"1,5,6,2,3,4",
		"1,5,6,2,4,3",
		"1,5,6,3,2,4",
		"1,5,6,3,4,2",
		"1,5,6,4,2,3",
		"1,5,6,4,3,2",
		"1,6,2,3,4,5",
		"1,6,2,3,5,4",
		"1,6,2,4,3,5",
		"1,6,2,4,5,3",
		"1,6,2,5,3,4",
		"1,6,2,5,4,3",
		"1,6,3,2,4,5",
		"1,6,3,2,5,4",
		"1,6,3,4,2,5",
		"1,6,3,5,2,4",
		"1,6,3,5,4,2",
		"1,6,4,2,3,5",
		"1,6,4,2,5,3",
		"1,6,4,3,2,5",
		"1,6,4,3,5,2",
		"1,6,4,5,2,3",
		"1,6,4,5,3,2",
		"1,6,5,2,3,4",
		"1,6,5,2,4,3",
		"1,6,5,3,2,4",
		"1,6,5,3,4,2",
		"1,6,5,4,2,3",
		"1,6,5,4,3,2",
		"1,6,3,4,5,2",
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

func SendEmail(email string, token string, name string, route string) error {
	clickurl := "http://localhost:8080/api/verifyUser?email=" + email +"&token=" + token
	data := []byte (`{"toEmail":` + `"` + email + `",` + `"url": "`+ clickurl +`",`+ `"name":"`+ name +`"}`)
	/*
	message := map[string]interface{}{
		"toEmail": email,
		"url": url,
		"name": name,
	}

	bytesRepresentation, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
		return err
	}
	*/
	req, err:= http.NewRequest("POST", route, bytes.NewBuffer(data))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "thehawkiswatching12345")

	if err != nil {
		return err
	}

	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}

/*
func SendEmail () {
	url := "https://mail.iecsemanipal.com/hawkeye/forgotpassword"

	payload := strings.NewReader("{\n    \"toEmail\": \"anshita_b@yahoo.co.in\",\n    \"url\": \"asdasd\",\n    \"name\": \"aasdasd\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "thehawkiswatching12345")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
*/
func LogRequest(r *http.Request, status string, err string) {

	body, _ := ioutil.ReadAll(r.Body)
	logInfo := LogInfo{
		Timestamp: time.Now(),
		Method:    r.Method,
		URL:       r.URL.Path,
		Body:      string(body),
		User:      r.Context().Value("User"),
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	log.Printf("%s\nTimestamp: %s\nMethod: %s\nURL: %s\nBody: %s\nUser: %v\nError: %s\n\n", status,
		logInfo.Timestamp.Format("Mon Jan _2 15:04:05 2006"), logInfo.Method, logInfo.URL, logInfo.Body, logInfo.User,
		err)
}
