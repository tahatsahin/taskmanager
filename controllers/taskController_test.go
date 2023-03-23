package controllers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"taskmanager/common"
	"testing"
)

func TestCreateTask(t *testing.T) {
	common.StartUp()
	// TODO: change this to abs path aswell
	file, err := os.Open("testTaskData.json")
	if err != nil {
		log.Fatalf("cannot open test file %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("cannot close file")
		}
	}(file)
	req, err := http.NewRequest(http.MethodPost, "/tasks", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

func TestGetTasks(t *testing.T) {
	common.StartUp()

	req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTasks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

func TestUpdateTask(t *testing.T) {
	common.StartUp()

	file, err := os.Open("testTaskData.json")
	if err != nil {
		log.Fatalf("cannot open file %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("cannot close file")
		}
	}(file)

	req, err := http.NewRequest(http.MethodPut, "/tasks/", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateTask)
	var endpointVar = map[string]string{
		"id": "63111153a5643c560f030d15",
	}
	req = mux.SetURLVars(req, endpointVar)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

func TestGetTaskById(t *testing.T) {
	common.StartUp()

	req, err := http.NewRequest(http.MethodGet, "/tasks/", nil)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTaskById)
	endpointVar := map[string]string{
		"id": "63111153a5643c560f030d15",
	}
	req = mux.SetURLVars(req, endpointVar)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDeleteTask(t *testing.T) {
	common.StartUp()

	req, err := http.NewRequest(http.MethodDelete, "/tasks/", nil)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)
	endpointVar := map[string]string{
		"id": "63111153a5643c560f030d15",
	}
	req = mux.SetURLVars(req, endpointVar)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
