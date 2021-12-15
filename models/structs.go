package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Username    string `json:"username" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type User struct {
	ID          primitive.ObjectID `bson:"id"`
	Username    string             `bson:"username"`
	DisplayName string             `bson:"displayName"`
	Password    string             `bson:"password"`
}
