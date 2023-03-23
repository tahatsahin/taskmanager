package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"taskmanager/common"
	"testing"
)

type Data struct {
	CreatedBy   string `json:"createdBy"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Task struct {
	Data Data `json:"data"`
}

func TestCreateTask(t *testing.T) {
	common.StartUp()

	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("cannot open file %v", err)
	}
	defer file.Close()

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

	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("cannot open file %v", err)
	}
	defer file.Close()

	req, err := http.NewRequest(http.MethodPut, "/tasks/63fddf39a5643c560f030d14", file)
	if err != nil {
		log.Fatalf("cannot create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateTask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
	log.Println(rr.Body)
}
