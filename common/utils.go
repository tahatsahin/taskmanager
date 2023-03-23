package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// define a configuration struct to store the variables in .env file
type configuration struct {
	SERVER, HOST, DBUSER, PWD, URI, Database, DBURI string
}

// define an error struct
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

// initConfig loads the .env file
func initConfig() {
	absPath, err := filepath.Abs("common/.env")
	absPath = strings.Split(absPath, "taskmanager")[0] + "taskmanager\\common\\.env"
	if err != nil {
		log.Fatalf("cannot get absolute path: %v", err)
	}

	err = godotenv.Load(os.ExpandEnv(absPath))
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

// DisplayAppError creates a standard error message for HTTP responses
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
