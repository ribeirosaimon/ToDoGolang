package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"learnGo/src/database"
	"learnGo/src/models"
	"learnGo/src/payloads"
	"time"
)

const (
	taskCollection    = "Task"
	contextRepository = 2
)

func FindAllTasks() (models.Tasks, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), contextRepository*time.Second)
	defer cancel()

	mongo := database.ConnectionMongo(taskCollection)
	cursor, err := mongo.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	list := models.Tasks{}

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return list, errors.New("erro na leitura")
		}

		list = append(list, task)
	}
	if err != nil {
		return list, errors.New("error find Tasks")
	}

	return list, nil
}

func FindTask(id string) (models.Task, error) {
	var err error
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New(fmt.Sprintf("Incorrect Id: %s", primitiveId))
	}
	ctx, cancel := context.WithTimeout(context.Background(), contextRepository*time.Second)
	defer cancel()

	mongo := database.ConnectionMongo(taskCollection)
	var result models.Task
	err = mongo.FindOne(ctx, bson.D{{"_id", primitiveId}}).Decode(&result)

	if err != nil {
		return models.Task{}, errors.New(fmt.Sprintf("Task with Id: %s not found", id))
	}
	return result, nil
}

func SaveTask(task payloads.TaskDTO) (models.Task, error) {

	ctx, cancel := context.WithTimeout(context.Background(), contextRepository*time.Second)
	defer cancel()

	newTask := models.Task{}
	if validateTask(task.Description) {
		return newTask, errors.New("task cannot be Null")
	}

	newTask.ToDoName = task.Description
	newTask.IsCompleted = false
	newTask.CreatedAt, newTask.UpdatedAt = time.Now(), time.Now()

	mongo := database.ConnectionMongo(taskCollection)

	res, err := mongo.InsertOne(ctx, newTask)
	if err != nil {
		return models.Task{}, err
	} else {
		newTask.ID = res.InsertedID.(primitive.ObjectID)
		return newTask, nil
	}
}

func EditTask(id string, task payloads.TaskDTO) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextRepository*time.Second)
	defer cancel()
	newTask := models.Task{}

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return newTask, errors.New(fmt.Sprintf("Task with Id: %s not found", id))
	}
	if validateTask(task.Description) {
		return newTask, errors.New("task cannot be Null")
	}
	mongo := database.ConnectionMongo(taskCollection)
	filter := bson.M{"_id": primitiveId}
	update := bson.M{"$set": bson.M{"task": task.Description}}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	result := mongo.FindOneAndUpdate(ctx, filter, update, &opt)

	if result.Err() != nil {
		return newTask, nil
	}

	decodeErr := result.Decode(&newTask)
	return newTask, decodeErr
}

func DeleteTask(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextRepository*time.Second)
	defer cancel()

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Task with Id: %s not found", id))
	}
	mongo := database.ConnectionMongo(taskCollection)
	isDeleted, err := mongo.DeleteOne(ctx, bson.D{
		{"_id", primitiveId},
	})

	if isDeleted.DeletedCount == 0 && err != nil {
		return false, errors.New("error to delete task")
	}
	return true, nil
}

func validateTask(dto string) bool {
	if dto == "" {
		return true
	}
	return false
}
