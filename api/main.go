package main

import (
	"os"

	"github.com/kelsier27/hawkeye19/api/app"
)

func main() {
	hawk := app.App{}
	hawk.Initialise()
	hawk.Run(os.Args)
}
