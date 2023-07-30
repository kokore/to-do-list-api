package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	toDoList "github.com/hugeman/todolist/handler/to_do_list"
	"github.com/hugeman/todolist/internal/config"
	"github.com/hugeman/todolist/internal/logz"
	"github.com/stretchr/testify/assert"
)

func TestGetToDoListSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListSearchSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?search=a", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderTitleASCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=title&orderType=ASC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderTitleDESCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=title&orderType=DESC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderDateASCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=Date&orderType=ASC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderDateDESCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=date&orderType=DESC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderStatusASCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=Date&orderType=ASC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestGetToDoListOrderStatusDESCSuccess(t *testing.T) {
	r := gin.Default()
	r.GET("/api/v1/to-do-list", toDoList.GetToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/api/v1/to-do-list?orderBy=date&orderType=DESC", nil)
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")

	expectedResponse := "{\"code\":0,\"message\":\"success\",\"data\":[{\"id\":\"64c64697708e2f1af13d506c\",\"title\":\"d\",\"description\":\"d\",\"date\":\"2023-07-30T11:53:50Z\",\"image\":\"d\",\"status\":\"IN_PROGRESS\"},{\"id\":\"64c4948f752146cb67c3f298\",\"title\":\"c\",\"description\":\"c\",\"date\":\"2023-07-30T10:48:51Z\",\"image\":\"c\",\"status\":\"COMPLETED\"},{\"id\":\"64c4942a25f3ffc4b55a642b\",\"title\":\"b\",\"description\":\"b\",\"date\":\"2023-07-30T10:43:10Z\",\"image\":\"b\",\"status\":\"COMPLETED\"},{\"id\":\"64c62fb18b71008502168dd6\",\"title\":\"a\",\"description\":\"a\",\"date\":\"2023-07-30T09:38:57Z\",\"image\":\"a\",\"status\":\"IN_PROGRESS\"}],\"statusCode\":200}"
	assert.Equal(t, expectedResponse, recorder.Body.String(), "Response body should match the expected response")
}

func TestCreateToDoListSuccess(t *testing.T) {
	r := gin.Default()
	r.POST("/api/v1/to-do-list", toDoList.CreateToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	requestBody := []byte(`{ "title": "d", "description": "d", "image": "d", "status": "IN_PROGRESS" }`)

	req, err := http.NewRequest("POST", "/api/v1/to-do-list", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")
}

func TestCreateToDoListFail(t *testing.T) {
	r := gin.Default()
	r.POST("/api/v1/to-do-list", toDoList.CreateToDoList)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	requestBody := []byte(`{ ""description": "d", "image": "d", "status": "IN_PROGRESS" }`)

	req, err := http.NewRequest("POST", "/api/v1/to-do-list", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Response status code should be 400")
}

func TestUpdateToDoListByIdSuccess(t *testing.T) {
	r := gin.Default()
	r.PUT("/api/v1/to-do-list/:id", toDoList.UpdateToDoListById)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	requestBody := []byte(`{ "title": "d", "description": "d", "image": "d", "status": "IN_PROGRESS" }`)

	req, err := http.NewRequest("PUT", "/api/v1/to-do-list/64c64697708e2f1af13d506c", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Response status code should be 200")
}

func TestUpdateToDoListByIdBodyFail(t *testing.T) {
	r := gin.Default()
	r.PUT("/api/v1/to-do-list/:id", toDoList.UpdateToDoListById)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	requestBody := []byte(`{ ""description": "d", "image": "d", "status": "IN_PROGRESS" }`)

	req, err := http.NewRequest("PUT", "/api/v1/to-do-list/64c64697708e2f1af13d506c", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Response status code should be 200")
}

func TestUpdateToDoListByIdIdMissingFail(t *testing.T) {
	r := gin.Default()
	r.PUT("/api/v1/to-do-list/:id", toDoList.UpdateToDoListById)

	os.Setenv("TZ", "Asia/Bangkok")

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}

	requestBody := []byte(`{ "title": "d", "description": "d", "image": "d", "status": "IN_PROGRESS" }`)

	req, err := http.NewRequest("PUT", "/api/v1/to-do-list/123", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "Error creating request")

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, "Response status code should be 200")
}
