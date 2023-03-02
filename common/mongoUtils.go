package common

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var session mongo.Session

// GetSession returns existing session, if there isn't any, it creates one
func GetSession() mongo.Session {
	if session == nil {
		createDbSession()
	}
	return session
}

// createDbSession creates a database session with given Database uri given in .env file
func createDbSession() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(AppConfig.DBURI).SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("[createDbConnection]: %v", err)
	}
	session, err = client.StartSession()
	if err != nil {
		log.Fatalf("[createDbSession]: %v", err)
	}
}

// addIndexes adds indexes to mongodb database
func addIndexes() {
	var err error
	// define indexes
	userIndex := mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: nil,
	}
	taskIndex := mongo.IndexModel{
		Keys: bson.M{
			"createdby": 1,
		},
		Options: nil,
	}
	// get a session
	session := GetSession()
	defer session.EndSession(context.TODO())

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// get collections
	userCol := session.Client().Database(AppConfig.Database).Collection("users")
	taskCol := session.Client().Database(AppConfig.Database).Collection("tasks")

	// create indexes on mongodb
	_, err = userCol.Indexes().CreateOne(ctx, userIndex)
	if err != nil {
		log.Fatalf("[userIndexCreate]: %v", err)
	}
	_, err = taskCol.Indexes().CreateOne(ctx, taskIndex)
	if err != nil {
		log.Fatalf("[taskIndexCreate]: %v", err)
	}

}
