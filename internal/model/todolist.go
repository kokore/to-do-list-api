package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusType string

type ToDoListSearchFilter struct {
}

const (
	IN_PROGRESS StatusType = "IN_PROGRESS"
	COMPLETED   StatusType = "COMPLETED"
)

func (s StatusType) IsValid() error {
	if s != IN_PROGRESS && s != COMPLETED {
		return errors.New("invalid status")
	}
	return nil
}

func DateFormat() *string {
	now := time.Now().Format(time.RFC3339)
	return &now
}

type ToDoList struct {
	Title       *string     `bson:"title"`
	Description string      `bson:"description"`
	Date        *string     `bson:"date"`
	Image       string      `bson:"image"`
	Status      *StatusType `bson:"status"`
}

type ToDoListDTO struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       *string            `json:"title"`
	Description string             `json:"description"`
	Date        *time.Time         `json:"date"`
	Image       string             `json:"image"`
	Status      *StatusType        `json:"status"`
}

type ToDoListById struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
}

type UpdateToDoList struct {
	Set ToDoList `bson:"$set"`
}

func CreateToDoListModel() ToDoList {
	return ToDoList{
		Date: DateFormat(),
	}
}

func UpdateToDoListModel(todolist ToDoList) UpdateToDoList {
	return UpdateToDoList{
		Set: todolist,
	}
}
