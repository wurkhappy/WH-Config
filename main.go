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
var EmailBroker string
var BalancedUsername string
var WebServer string
var MDPBroker string
var RMQBroker string
var WebAppRedis string
var TransactionsBroker string
var TransactionsQueue string
var TransactionsExchange string
var PDFService string

type Configuration struct {
	UserService          string
	PaymentInfoService   string
	AgreementsService    string
	CommentsService      string
	EmailExchange        string
	EmailQueue           string
	EmailBroker          string
	BalancedUsername     string
	WebServer            string
	MDPBroker            string
	RMQBroker            string
	WebAppRedis          string
	TransactionsBroker   string
	TransactionsQueue    string
	TransactionsExchange string
	PDFService           string
}

func Prod() {
	file, err := os.Open("/home/wh//WH-Config/config.json")
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
	EmailBroker = configuration.EmailBroker
	BalancedUsername = configuration.BalancedUsername
	WebServer = configuration.WebServer
	MDPBroker = configuration.MDPBroker
	RMQBroker = configuration.RMQBroker
	WebAppRedis = configuration.WebAppRedis
	TransactionsBroker = configuration.TransactionsBroker
	TransactionsQueue = configuration.TransactionsQueue
	TransactionsExchange = configuration.TransactionsExchange
	PDFService = configuration.PDFService
}
