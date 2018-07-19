package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Server, DBUser, DBPWd, Database string
}

var AppConfig configuration

func StartUp() {
	//initConfig()

	///

}

func initConfig() {
	loadAPpConfig()
}

func loadAPpConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadingConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}