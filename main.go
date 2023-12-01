package main

import (
	"fmt"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handlers.RootHandler)
	v1.POST("/tasks", taskHandler.PostTasksHandler)
	v1.GET("/tasks/:id", taskHandler.GetTaskByIDHandler)
	v1.POST("/tasks/:id", taskHandler.PostSubTaskByID)
	v1.GET("tasks/filter", taskHandler.FilterTaskHandler)

	router.Run()

}

// type Task struct {
// 	gorm.Model
// 	Title       string
// 	Description string
// 	ParentID    *uint
// 	Children    []Task `gorm:"foreignKey:ParentID"`
// }

// func main() {
// 	//Connect to the SQLite database
// 	dsn := "host=localhost user=root password=secret dbname=DBGeneric port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// Auto Migrate the schema
// 	db.AutoMigrate(&Task{})

// 	parentTask := Task{Title: "Parent Task", Description: "This is the parent task"}
// 	childTask := Task{Title: "Child Task", Description: "This is a child task"}

// 	// Link the child task to the parent task
// 	parentTask.Children = append(parentTask.Children, childTask)

// 	// Save the parent task to the database
// 	db.Create(&parentTask)

// }
