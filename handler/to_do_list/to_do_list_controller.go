package to_do_list

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugeman/todolist/internal/logz"
	"github.com/hugeman/todolist/internal/model"
	"github.com/hugeman/todolist/internal/request/queryParams"
	"github.com/hugeman/todolist/internal/response"
	"github.com/hugeman/todolist/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type ToDoListRequest struct {
	Title       *string           `json:"title"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	Status      *model.StatusType `json:"status"`
}

func (req ToDoListRequest) validate() map[string]string {
	errMap := map[string]string{}

	if utils.IsEmptyString(req.Title) {
		errMap["title"] = "Please fill title"
	}
	if err := req.Status.IsValid(); err != nil {
		errMap["status"] = err.Error()
	}

	return errMap
}

type GetToDoListRequest struct {
	OrderQueryParam  queryParams.OrderQueryParam
	SearchQueryParam queryParams.SearchQueryParam
}

func GetToDoList(c *gin.Context) {
	logz.Logger.Info("Start Get To Do List")
	startTime := time.Now()

	var params GetToDoListRequest
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.InvalidQueryParams, http.StatusBadRequest, err.Error()))
		return
	}

	result, err := FindAllToDoListToDB(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(response.UnableToFindToDoList, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.OK(result))
	logz.Logger.Info("Get To Do List Success", zap.Int64("duration", time.Since(startTime).Milliseconds()))
}

func CreateToDoList(c *gin.Context) {
	logz.Logger.Info("Start Create To Do List")
	startTime := time.Now()
	var req ToDoListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}
	if errMap := req.validate(); len(errMap) != 0 {
		c.JSON(http.StatusBadRequest, response.ErrField(response.InvalidRequestJSONString, http.StatusBadRequest, "Please complete the information.", errMap))
		return
	}

	result, err := CreateToDoListToDB(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(response.UnableToCreateToDoList, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.OK(result))
	logz.Logger.Info("Create To Do List Success", zap.Int64("duration", time.Since(startTime).Milliseconds()))
}

type ToDoListIdRequest struct {
	ID string `uri:"id" binding:"required"`
}

func UpdateToDoListById(c *gin.Context) {
	logz.Logger.Info("Start Update To Do List")
	startTime := time.Now()

	var params ToDoListIdRequest
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusNotFound, ""))
		return
	}

	idToUpdate, err := primitive.ObjectIDFromHex(params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusNotFound, ""))
		return
	}

	var req ToDoListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}
	if errMap := req.validate(); len(errMap) != 0 {
		c.JSON(http.StatusBadRequest, response.ErrField(response.InvalidRequestJSONString, http.StatusBadRequest, "Please complete the information.", errMap))
		return
	}

	err = UpdateToDoListToDB(idToUpdate, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Err(response.UnableToUpdateToDoList, http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.OK(""))
	logz.Logger.Info("Update To Do List Success", zap.Int64("duration", time.Since(startTime).Milliseconds()))

}
