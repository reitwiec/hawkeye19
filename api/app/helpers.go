package app

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	SimilarStringRatio = 0.65
	CorrectAnswer      = 1
	CloseAnswer        = 2
	IncorrectAnswer    = 3
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
	if ratio >= SimilarStringRatio {
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

func SanitizeAnswer(s string) string {
	s = Sanitize(s)
	//remove special characters
	reg, err := regexp.Compile("[^a-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")
	return processedString
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
	permutations := [2]string{
		"0,2,3,4,5",
		//"0,2,4,3,5",
		"0,3,2,4,5",
		//"0,3,4,2,5",
		//"0,4,2,3,5",
		//"0,4,3,2,5",
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(2)
	return permutations[n]
}

func GetNextRegion(unlockOrder string) uint8 {
	//traverse till non 0 number is encountered

	l := len(unlockOrder)
	for i := 0; i < l; i++ {
		if unlockOrder[i] != ',' && unlockOrder[i] != '0' {
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
func SendFPEmail(email string, token string, name string) error {
	clickurl := "https://hawkeye.iecsemanipal.com/api/forgotPassword?email=" + email + "&token=" + token
	data := []byte(`{"toEmail":` + `"` + email + `",` + `"url": "` + clickurl + `",` + `"name":"` + name + `"}`)
	req, err := http.NewRequest("POST", "https://mail.iecsemanipal.com/hawkeye/forgotpassword", bytes.NewBuffer(data))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "thehawkiswatching12345")

	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}

func SendVUEmail(email string, token string, name string) error {
	clickurl := "https://hawkeye.iecsemanipal.com/api/verifyUser?email=" + email + "&token=" + token
	data := []byte(`{"toEmail":` + `"` + email + `",` + `"url": "` + clickurl + `",` + `"name":"` + name + `"}`)
	req, err := http.NewRequest("POST", "https://mail.iecsemanipal.com/hawkeye/emailverification", bytes.NewBuffer(data))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "thehawkiswatching12345")

	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
func LogRequest(r *http.Request, status string, err string) {

	body, _ := ioutil.ReadAll(r.Body)
	logInfo := LogInfo{
		Timestamp: time.Now(),
		Method:    r.Method,
		URL:       r.URL.Path,
		//@TODO add body after verified
		//Body:      string(body),
		Body: "",
		User: r.Context().Value("User"),
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	log.Printf("%s\nTimestamp: %s\nMethod: %s\nURL: %s\nBody: %s\nUser: %v\nError: %s\n\n", status,
		logInfo.Timestamp.Format("Mon Jan _2 15:04:05 2006"), logInfo.Method, logInfo.URL, logInfo.Body, logInfo.User,
		err)
}

func (hawk *App) GetUser(w http.ResponseWriter, r *http.Request) {
	currUser, err := GetCurrUser(r)
	if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Could not read cookie data", nil, http.StatusInternalServerError, w)
		return
	}
	user := User{}
	err = hawk.DB.Where("username = ?", currUser.Username).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		ResponseWriter(false, "User not registered", nil, http.StatusOK, w)
		return
	} else if err != nil {
		LogRequest(r, ERROR, err.Error())
		ResponseWriter(false, "Database error", nil, http.StatusInternalServerError, w)
		return
	}
	user.Password = ""
	ResponseWriter(true, "User fetched", user, http.StatusOK, w)
}

//@TODO
//func SanitizeAnswer
