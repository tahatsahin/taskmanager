package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"taskmanager/models"
)

type UserRepository struct {
	C *mongo.Collection
}

// CreateUser creates a user with given user model
func (r *UserRepository) CreateUser(user models.User) error {
	objId := primitive.NewObjectID()
	user.Id = objId
	hPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hPass
	// clear the incoming text password
	user.Password = ""
	_, err = r.C.InsertOne(context.TODO(), &user)
	return err

}

// Login logs user in the system
func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	var res models.User
	err = r.C.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&res)
	if err != nil {
		log.Fatalf("there is no such user, check your confidentials...: %v", err)
	}

	err = bcrypt.CompareHashAndPassword(res.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
		return u, err
	}
	return res, nil
}
