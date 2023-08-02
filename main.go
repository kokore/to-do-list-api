package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	toDoList "github.com/hugeman/todolist/handler/to_do_list"
	"github.com/hugeman/todolist/internal/config"
	"github.com/hugeman/todolist/internal/logz"
)

func main() {
	initial()
	ginEngine := getGinEngine()

	port := fmt.Sprintf(":%s", config.Config.App.Port)
	ginEngine.Run(port)
}

func initial() {
	os.Setenv("TZ", "Asia/Bangkok")

	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)

	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := logz.Init(); err != nil {
		log.Fatal(err)
	}
}

func getGinEngine() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.CORS())
	router.Use(cors.Default())

	root := router.Group("/api/v1")

	toDoListRouter := root.Group("/to-do-list")

	toDoListRouter.GET("", toDoList.GetToDoList)
	toDoListRouter.POST("", toDoList.CreateToDoList)
	toDoListRouter.PUT("/:id", toDoList.UpdateToDoListById)

	return router
}
