package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"taskmanager/common"
)

// Context struct used for maintaining HTTP request context
type Context struct {
	MongoSession mongo.Session
}

// Close session
func (c *Context) Close() {
	c.MongoSession.EndSession(context.TODO())
}

// DbCollection returns the collection with given name
func (c *Context) DbCollection(name string) *mongo.Collection {
	return c.MongoSession.Client().Database(common.AppConfig.Database).Collection(name)
}

// NewContext creates context for mongodb session
func NewContext() *Context {
	session := common.GetSession()
	ctx := &Context{
		MongoSession: session,
	}
	return ctx
}
