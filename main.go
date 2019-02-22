package main

import (
	"github.com/kelsier27/hawkeye19/app"
	"os"
)

func main (){
	hawk:=app.App {}
	hawk.Initialise ()
	hawk.Run (os.Args)
}

