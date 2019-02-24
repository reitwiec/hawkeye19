package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gorilla/securecookie"
)

func main() {
	s := fmt.Sprintf(`HAWK_HASH_KEY=%s
HAWK_BLOCK_KEY=%s
`, securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

	err := ioutil.WriteFile(".env.keys", []byte(s), 0644)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	fmt.Println("Wrote keys to .env.keys")
}
