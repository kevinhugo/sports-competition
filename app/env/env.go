package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Service Setting
var SERVICE_HOST string
var SERVICE_PORT string
var SERVICE_ROOT_PATH string

// \ Service Setting

// Database Setting
var DB_USER string
var DB_PASS string
var DB_NAME string
var DB_HOST string
var DB_PORT string
var DB_PING_DELAY_IN_SECONDS int

// \ Database Setting

// Logger Configuration
var LOGGER_SHOW_DEBUG bool
var LOGGER_SHOW_INFO bool
var LOGGER_SHOW_WARNING bool
var LOGGER_SHOW_ERROR bool

// \ Logger Configuration

// Token Configuration
var ACCESS_TOKEN_EXPIRED_IN_MINUTES int
var ACCESS_TOKEN_SECRET_KEY []byte

// \ Token Configuration

// Outbound Config
var OUTBOUND_CONNECTION_TIMEOUT_IN_SECOND int
var INBOUND_CONNECTION_TIMEOUT_IN_SECOND int

// \ Outbound Config

func init() {
	LoadEnv()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
		return
	}
	initEnvVariables()
	abortInitializationOnMissingEnv()
	log.Println("Success load .env file")
	log.Println(SERVICE_HOST)
}

func initEnvVariables() {
	setEnv(&SERVICE_HOST, "SERVICE_HOST")
	setEnv(&SERVICE_PORT, "SERVICE_PORT")
	setEnv(&SERVICE_ROOT_PATH, "SERVICE_ROOT_PATH")

	setEnv(&DB_USER, "DB_USER")
	setEnv(&DB_PASS, "DB_PASS")
	setEnv(&DB_NAME, "DB_NAME")
	setEnv(&DB_HOST, "DB_HOST")
	setEnv(&DB_PORT, "DB_PORT")
	setEnv(&DB_PING_DELAY_IN_SECONDS, "DB_PING_DELAY_IN_SECONDS", 60)

	setEnv(&ACCESS_TOKEN_SECRET_KEY, "ACCESS_TOKEN_SECRET_KEY")
	setEnv(&ACCESS_TOKEN_EXPIRED_IN_MINUTES, "ACCESS_TOKEN_EXPIRED_IN_MINUTES")

	setEnv(&LOGGER_SHOW_DEBUG, "LOGGER_SHOW_DEBUG", true)
	setEnv(&LOGGER_SHOW_INFO, "LOGGER_SHOW_INFO", true)
	setEnv(&LOGGER_SHOW_WARNING, "LOGGER_SHOW_WARNING", true)
	setEnv(&LOGGER_SHOW_ERROR, "LOGGER_SHOW_ERROR", true)

	setEnv(&OUTBOUND_CONNECTION_TIMEOUT_IN_SECOND, "OUTBOUND_CONNECTION_TIMEOUT_IN_SECOND", 10)
	setEnv(&INBOUND_CONNECTION_TIMEOUT_IN_SECOND, "INBOUND_CONNECTION_TIMEOUT_IN_SECOND", 10)
}

func abortInitializationOnMissingEnv() {
	if SERVICE_ROOT_PATH == "" {
		panic("Missing critical SERVICE_ROOT_PATH in .env .")
	}
}

/*
Assign given envVariable based on result of os.Getenv(envname), returns defaultValue if give os.Getenv(envname) result is empty string.

envVariable supports *string, *[]byte, *int, and *bool
*/
func setEnv(envVariable interface{}, envName string, defaultValue ...interface{}) {
	var envValue = os.Getenv(envName)
	switch envVariable.(type) {
	case *string:
		*envVariable.(*string) = envValue
		if *envVariable.(*string) == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*string) = defaultValue[0].(string)
			}
		}
	case *[]byte:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*[]byte) = defaultValue[0].([]byte)
			} else {
				*envVariable.(*[]byte) = []byte("")
			}
		} else {
			*envVariable.(*[]byte) = []byte(envValue)
		}
	case *int:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*int) = defaultValue[0].(int)
			} else {
				*envVariable.(*int) = 0
			}
		} else {
			newInt, _ := strconv.Atoi(os.Getenv(envName))
			*envVariable.(*int) = newInt
		}
	case *bool:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*bool) = defaultValue[0].(bool)
			} else {
				*envVariable.(*bool) = false
			}
		} else {
			if envValue == "1" || envValue == "true" {
				*envVariable.(*bool) = true
			} else {
				*envVariable.(*bool) = false
			}
		}
	default:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*interface{}) = defaultValue[0]
			}
		} else {
			*envVariable.(*interface{}) = envValue
		}
	}
}
