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
var PaymentsService string
var CommentsService string
var EmailQueue string
var BalancedUsername string
var WebServer string
var MDPBroker string
var RMQBroker string
var WebAppRedis string
var TransactionsQueue string
var PDFService string
var TasksService string
var MainExchange string
var DeadLetterExchange string
var PDFTemplatesService string

type Configuration struct {
	UserService         string
	PaymentInfoService  string
	AgreementsService   string
	PaymentsService     string
	CommentsService     string
	EmailQueue          string
	BalancedUsername    string
	WebServer           string
	MDPBroker           string
	RMQBroker           string
	WebAppRedis         string
	TransactionsQueue   string
	PDFService          string
	TasksService        string
	MainExchange        string
	DeadLetterExchange  string
	PDFTemplatesService string
}

func Prod() {
	file, err := os.Open("/home/wh/WH-Config/config.json")
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
	PaymentsService = configuration.PaymentsService
	CommentsService = configuration.CommentsService
	EmailQueue = configuration.EmailQueue
	BalancedUsername = configuration.BalancedUsername
	WebServer = configuration.WebServer
	MDPBroker = configuration.MDPBroker
	RMQBroker = configuration.RMQBroker
	WebAppRedis = configuration.WebAppRedis
	TransactionsQueue = configuration.TransactionsQueue
	PDFService = configuration.PDFService
	TasksService = configuration.TasksService
	MainExchange = configuration.MainExchange
	DeadLetterExchange = configuration.DeadLetterExchange
	PDFTemplatesService = configuration.PDFTemplatesService
}
