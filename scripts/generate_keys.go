package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func generateKey(size int) string {
	rand.Seed(time.Now().UnixNano())
	s := "1234567890qwertyuiopasdfghjklzxcvbnmMNBVCXZLKJHFDSAPOIUYTREWQ"
	l := len(s)
	var j int
	s1 := ""
	for i := 0; i < size; i++ {
		j = rand.Intn(l)
		s1 += string(s[j])
	}
	return (s1)

}

func main() {
	hashKey := generateKey(32)
	blockKey := generateKey(32)
	s := fmt.Sprintf(`HAWK_HASH_KEY=%s
HAWK_BLOCK_KEY=%s
`, hashKey, blockKey)

	err := ioutil.WriteFile(".env.keys", []byte(s), 0644)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Println("Wrote keys to .env.keys")
}
