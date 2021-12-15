package database

import (
	"context"
	"log"

	"github.com/pierorolando1/crudmongo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var collection *mongo.Collection
var ctx = context.TODO()

func client() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://64.225.29.144:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client

}

var Collection = client().Database("crudmongo").Collection("users")

func FindByUsername(username string) (*models.User, error) {
	var user *models.User
	err := Collection.FindOne(ctx, map[string]interface{}{"username": username}).Decode(&user)
	return user, err
}

func CreateUser(user interface{}) (interface{}, error) {
	doc, err := Collection.InsertOne(ctx, user)
	return doc, err
}

// get user by username
// func FindByUsername(username string) (interface{}, error) {
