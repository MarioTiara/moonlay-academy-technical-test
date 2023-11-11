package main

import (
	"fmt"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/gin-gonic/gin"
	priority "github.com/marioTiara/todolistwebapi/Priority"
	task "github.com/marioTiara/todolistwebapi/Task"
	"github.com/marioTiara/todolistwebapi/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	priorityRepository := priority.NewRepository(db)
	priorityService := priority.NewService(priorityRepository)
	priorityHandler := handlers.NewPriorityHandler(priorityService)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handlers.RootHandler)
	v1.POST("/tasks", taskHandler.PostTasksHandler)
	v1.POST("/priorities", priorityHandler.PostPriorityHandler)

	router.Run()

}
