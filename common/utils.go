package common

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type configuration struct {
	SERVER, HOST, DBUSER, PWD, URI, Database, DBURI string
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

// AppConfig holds the configuration values from .env file
var AppConfig configuration

func initConfig() {
	err := godotenv.Load("common/.env")
	if err != nil {
		panic(err)
	}
	loadAppConfig()
}

// loadAppConfig reads .env and places into AppConfig
func loadAppConfig() {
	AppConfig = configuration{
		SERVER:   os.Getenv("SERVER"),
		HOST:     os.Getenv("HOST"),
		DBUSER:   os.Getenv("DBUSER"),
		PWD:      os.Getenv("PWD"),
		URI:      fmt.Sprintf("%v://%v:%v@%v", AppConfig.SERVER, AppConfig.DBUSER, AppConfig.PWD, AppConfig.HOST),
		Database: os.Getenv("DATABASE"),
		DBURI:    os.Getenv("DBURI"),
	}
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		_, err := w.Write(j)
		if err != nil {
			return
		}
	}
}
