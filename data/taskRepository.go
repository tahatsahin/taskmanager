package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"taskmanager/models"
	"time"
)

type TaskRepository struct {
	C *mongo.Collection
}

// CreateTask Create creates a new task
func (r *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	objId := primitive.NewObjectID()
	task.Id = objId
	task.CreatedOn = time.Now()
	task.Status = "created"
	_, err := r.C.InsertOne(context.TODO(), &task)
	if err != nil {
		log.Fatalf("cannot create task: %v", err)
	}
	return task, nil
}
