package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ToDoName    string             `bson:"task"`
	IsCompleted bool               `bson:"isCompleted"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

type Tasks []Task
