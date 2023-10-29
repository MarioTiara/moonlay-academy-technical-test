package main

import (
	"fmt"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/marioTiara/todolistwebapi/models"
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

	priority := models.Priority{ID: 1}

	db.First(&priority)

	fmt.Println(priority)

	task := models.Task{ID: 1}
	db.First(&task)
	fmt.Println(task)

	file := models.Files{ID: 1}
	db.First(&file)
	fmt.Println(file)

}
