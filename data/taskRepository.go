package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// UpdateTask updates a task by given id
func (r *TaskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	// partial update on mongodb
	updOpt := options.Update().SetUpsert(true)
	_, err := r.C.UpdateOne(context.TODO(), bson.M{"_id": task.Id},
		bson.M{"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		}}, updOpt)
	if err != nil {
		return &models.Task{}, err
	}
	return task, nil
}

// GetTasks returns all tasks
func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	cursor, err := r.C.Find(context.TODO(), bson.M{})
	var tasks []models.Task
	if err != nil {
		log.Fatalf("cannot retrieve tasks %v", err)
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var result models.Task
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, result)
	}
	return tasks, nil
}
