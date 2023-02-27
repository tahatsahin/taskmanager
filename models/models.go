package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	User struct {
		Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		FirstName    string             `json:"firstName"`
		LastName     string             `json:"lastName"`
		Email        string             `json:"email"`
		Password     string             `json:"password,omitempty"`
		HashPassword []byte             `json:"hashPassword,omitempty"`
	}
	Task struct {
		Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		CreatedBy   string             `json:"createdBy"`
		Name        string             `json:"name"`
		Description string             `json:"description"`
		CreatedOn   time.Time          `json:"createdOn,omitempty"`
		Due         time.Time          `json:"due,omitempty"`
		Status      string             `json:"status,omitempty"`
		Tags        []string           `json:"tags,omitempty"`
	}
)
