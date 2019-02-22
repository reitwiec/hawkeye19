package main

import (
	"app"
	"os"
)
func main (){
	hawk:=app.App {}
	hawk.Initialise ()
	hawk.Run (os.Args)
}
