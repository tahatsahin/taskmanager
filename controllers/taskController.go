package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
)

// CreateTask handler for HTTP POST - /tasks
// insert a new task doc
// TODO: Test Due
func CreateTask(w http.ResponseWriter, r *http.Request) {
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
	task := &dataResource.Data
	ctx := NewContext()
	defer ctx.Close()
	c := ctx.DbCollection("tasks")
	repo := &data.TaskRepository{C: c}
	// insert
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
