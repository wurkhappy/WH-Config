package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var UserService string
var PaymentInfoService string
var AgreementsService string
var CommentsService string
var EmailExchange string
var EmailQueue string
var EmailURI string

type Configuration struct {
	UserService        string
	PaymentInfoService string
	AgreementsService  string
	CommentsService    string
	EmailExchange      string
	EmailQueue         string
	EmailURI           string
}

func Prod() {
	absPath, _ := filepath.Abs("../WH-Config/config.json")
	log.Print("production config")
	_, err := os.Open(absPath)
	if err != nil {
		log.Printf("path err %s", err.Error())
	}

	// decoder := json.NewDecoder(file)
	// var configuration *Configuration
	// err = decoder.Decode(&configuration)
	// if err != nil {
	// 	log.Print(err)
	// }
	// setGlobalVars(configuration)
	// log.Print(configuration)
}

func Test() {
	log.Print("test config")
	file, _ := os.Open("test_config.json")
	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	decoder.Decode(&configuration)
	setGlobalVars(configuration)
}

func setGlobalVars(configuration *Configuration) {
	UserService = configuration.UserService
	PaymentInfoService = configuration.PaymentInfoService
	AgreementsService = configuration.AgreementsService
	CommentsService = configuration.CommentsService
	EmailExchange = configuration.EmailExchange
	EmailQueue = configuration.EmailQueue
	EmailURI = configuration.EmailURI
}
