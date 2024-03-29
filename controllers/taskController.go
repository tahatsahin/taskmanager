package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbCollectionTask = "tasks"

// CreateTask handler for HTTP POST - /tasks
// insert a new task doc
// TODO: Create task with a easier due date format like "DD.MM.YYYY"
func CreateTask(w http.ResponseWriter, r *http.Request) {
	// create a data variable to store
	var dataResource TaskResource
	// decode incoming json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid task data",
			500,
		)
		return
	}
	// take data into task
	task := &dataResource.Data
	// create context
	ctx := NewContext()
	defer ctx.Close()
	// get db collection
	c := ctx.DbCollection(dbCollectionTask)
	repo := &data.TaskRepository{C: c}
	// insert data
	task, err = repo.CreateTask(task)
	if err != nil {
		return
	}
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write(j)
		if err != nil {
			return
		}
	}
}

// UpdateTask handler for HTTP PUT - /tasks/{id}
// update an existing task doc
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// get id from request
	vars := mux.Vars(r)
	// convert hex id to objectID
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Fatalf("cannot get id from request: %v, %v", err, id)
	}
	var dataResource TaskResource
	// decode the incoming task json
	err = json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"invalid task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.Id = id
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionTask)
	repo := &data.TaskRepository{C: c}

	// update an existing task document
	if _, err := repo.UpdateTask(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// GetTasks handler for HTTP GET - /tasks
// returns all task docs
func GetTasks(w http.ResponseWriter, _ *http.Request) {
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionTask)
	repo := &data.TaskRepository{C: c}
	tasks, err := repo.GetTasks()
	if err != nil {
		log.Fatal(err)
	}
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		return
	}
}

// GetTaskById handler for HTTP GET - /tasks/{id}
// returns a single task doc by given id
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	// get id from request
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionTask)
	repo := &data.TaskRepository{C: c}
	task, err := repo.GetById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"an unexpected error has occurred",
				500,
			)
			return
		}
	}
	if j, err := json.Marshal(&task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(j)
		if err != nil {
			return
		}
	}
}

// DeleteTask handler for HTTP DELETE - /tasks/{id}
// deletes task by given id
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection(dbCollectionTask)
	repo := &data.TaskRepository{C: c}
	err := repo.DeleteById(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"an unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
