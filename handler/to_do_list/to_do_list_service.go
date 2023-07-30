package to_do_list

import (
	"github.com/hugeman/todolist/internal/model"
	"github.com/hugeman/todolist/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllToDoListToDB(params GetToDoListRequest) ([]model.ToDoListDTO, error) {
	filter := bson.M{}
	sort := bson.D{}

	if params.OrderQueryParam.GetOrderTypeByDefault() != "" {
		sort = bson.D{
			{Key: params.OrderQueryParam.OrderBy, Value: utils.SortConditions(params.OrderQueryParam.GetOrderTypeByDefault())},
		}
	}

	if params.SearchQueryParam.HasSearchText() {
		searchFilter := bson.M{
			"$or": []bson.M{
				{"title": bson.M{"$regex": params.SearchQueryParam.GetSearchText(), "$options": "i"}},
				{"description": bson.M{"$regex": params.SearchQueryParam.GetSearchText(), "$options": "i"}},
			},
		}

		for key, value := range searchFilter {
			filter[key] = value
		}
	}

	findAllResult, err := FindAllToDoList(filter, sort)
	if err != nil {
		return nil, err
	}
	return findAllResult, nil
}

func CreateToDoListToDB(insertToDoListRequest ToDoListRequest) (*string, error) {
	todoList := model.CreateToDoListModel()
	todoList.Title = insertToDoListRequest.Title
	todoList.Description = insertToDoListRequest.Description
	todoList.Image = insertToDoListRequest.Image
	todoList.Status = insertToDoListRequest.Status

	insertResult, err := InsertToDoListToDb(&todoList)
	if err != nil {
		return nil, err
	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return &insertedID, nil
}

func UpdateToDoListToDB(id primitive.ObjectID, updateToDoListRequest ToDoListRequest) error {
	todoList := model.CreateToDoListModel()
	todoList.Title = updateToDoListRequest.Title
	todoList.Description = updateToDoListRequest.Description
	todoList.Image = updateToDoListRequest.Image
	todoList.Status = updateToDoListRequest.Status

	filter := model.ToDoListById{
		ID: id,
	}

	updateToDoList := model.UpdateToDoListModel(todoList)

	err := UpdateToDoListToDb(filter, updateToDoList)
	if err != nil {
		return err
	}
	return nil
}
