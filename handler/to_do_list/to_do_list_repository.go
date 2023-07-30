package to_do_list

import (
	"context"
	"fmt"

	"github.com/hugeman/todolist/internal/db"
	"github.com/hugeman/todolist/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllToDoList(filter primitive.M, sort primitive.D) ([]model.ToDoListDTO, error) {
	if err := db.InitalDatabase(); err != nil {
		return nil, err
	}
	fmt.Println("connecting database...")

	ctx := db.Ctx
	var toDoListSlice []model.ToDoListDTO

	collection := db.Database.Collection("list")
	cursor, err := collection.Find(ctx, filter, options.Find().SetSort(sort))
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var todo model.ToDoListDTO
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		toDoListSlice = append(toDoListSlice, todo)
	}

	defer db.Cancel()
	defer cursor.Close(ctx)
	defer db.Client.Disconnect(ctx)
	fmt.Println("disconnect database")

	return toDoListSlice, nil
}

func FindToDoListById(id string) (*model.ToDoListDTO, error) {
	if err := db.InitalDatabase(); err != nil {
		return nil, err
	}
	fmt.Println("connecting database...")

	ctx := db.Ctx
	var result model.ToDoListDTO

	filter := bson.M{"_id": id}

	collection := db.Database.Collection("list")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	defer db.Cancel()
	defer db.Client.Disconnect(ctx)
	fmt.Println("disconnect database")

	return &result, nil
}

func InsertToDoListToDb(todolist *model.ToDoList) (*mongo.InsertOneResult, error) {
	if err := db.InitalDatabase(); err != nil {
		return nil, err
	}
	fmt.Println("connecting database...")

	ctx := db.Ctx
	collection := db.Database.Collection("list")
	result, err := collection.InsertOne(ctx, todolist)
	if err != nil {
		return nil, err
	}

	defer db.Cancel()
	defer db.Client.Disconnect(ctx)
	fmt.Println("disconnect database")

	return result, nil
}

func UpdateToDoListToDb(filter model.ToDoListById, update model.UpdateToDoList) error {
	if err := db.InitalDatabase(); err != nil {
		return err
	}
	fmt.Println("connecting database...")

	ctx := db.Ctx
	collection := db.Database.Collection("list")

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	defer db.Cancel()
	defer db.Client.Disconnect(ctx)
	fmt.Println("disconnect database")

	return nil
}
