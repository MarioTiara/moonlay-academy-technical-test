package main

import (
	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/gin-gonic/gin"
	"github.com/marioTiara/todolistwebapi/handlers"
)

func main() {

	// dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// taskRepository := task.NewRepository(db)

	// tasks, _ := taskRepository.FindAll()

	// fmt.Println(tasks)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handlers.RootHandler)

	router.Run()

}
