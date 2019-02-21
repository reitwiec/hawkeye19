package main

import (
	"github.com/anshitabaid/miniHawk/app"
	"os"
)
func main (){
	hawk:=app.App {}
	hawk.Initialise ()
	hawk.Run (os.Args)
}