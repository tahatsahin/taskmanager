package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"taskmanager/common"
	"testing"
)

func TestRegister(t *testing.T) {
	common.StartUp()
	// TODO: change this to abs path aswell
	file, err := os.Open("testUserRegisterData.json")
	if err != nil {
		log.Fatalf("cannot open test file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("cannot close file")
		}
	}(file)
	req, err := http.NewRequest(http.MethodPost, "/users/register", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Register)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

func TestLogin(t *testing.T) {
	common.StartUp()
	// TODO: change this to abs path aswell
	file, err := os.Open("testUserLoginData.json")
	if err != nil {
		log.Fatalf("cannot open test file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("cannot close file")
		}
	}(file)
	req, err := http.NewRequest(http.MethodPost, "/users/login", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDeleteUser(t *testing.T) {
	common.StartUp()
	// TODO: change this to abs path aswell
	file, err := os.Open("testUserLoginData.json")
	if err != nil {
		log.Fatalf("cannot open test file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("cannot close file")
		}
	}(file)
	req, err := http.NewRequest(http.MethodDelete, "/users", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
