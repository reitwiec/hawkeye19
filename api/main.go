package main

import (
	"os"

	"github.com/kelsier27/hawkeye19/api/app"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
	gotenv.Load(".env.keys")
}

func main() {
	hawk := app.App{}
	hawk.Initialise()
	hawk.Run(os.Args)
}
