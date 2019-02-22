package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func LoadDefaultConfiguration () {
	Configuration = Settings{
		ServerAddr: ":8080",
		DBUsername: "hawkadmin",
		DBPassword: "neverat12",
		DBName:     "hawkeyedb",
		HashKey: "",
		BlockKey: "",
	}
}

func ConfigureInstance () error {
	jsonConfig, err := ioutil.ReadFile("config.json")
	if err != nil {
		//@TODO: log the error
		fmt.Println ("File not read, switching to default/s/t" + err.Error ())
		LoadDefaultConfiguration ()
		return err
	}
	err = json.Unmarshal(jsonConfig, &Configuration)
	if err !=  nil {
		fmt.Println ("Error in unmarshalling config, switching to default\n\t" + err.Error ())
		LoadDefaultConfiguration()
		return err
	}
	fmt.Println ("Config file successfully read")
	return nil
}

