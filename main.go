package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	absPath, _ := filepath.Abs(".")
	directories := strings.Split(absPath, "/")
	file, err := os.Open("/" + directories[1] + "/" + directories[2] + "/WH-Config/config.json")
	if err != nil {
		log.Printf("path err %s", err.Error())
	}

	decoder := json.NewDecoder(file)
	var configuration *Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Print(err)
	}
	setGlobalVars(configuration)
}

func Test() {
	absPath, _ := filepath.Abs(".")
	directories := strings.Split(absPath, "/")
	file, err := os.Open("/" + directories[1] + "/" + directories[2] + "/WH-Config/test_config.json")
	if err != nil {
		log.Printf("path err %s", err.Error())
	}

	decoder := json.NewDecoder(file)
	var configuration *Configuration
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Print(err)
	}
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
