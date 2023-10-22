package main

import (
	"fmt"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/marioTiara/todolistwebapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define the models to generate migrations for.
var model = []any{
	// &models.Task{},
	// &models.Priority{},
	// &models.File{},
}

func main() {
	// stmts, err := gormschema.New("postgres").Load(model...)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
	// 	os.Exit(1)
	// }
	// io.WriteString(os.Stdout, stmts)

	dsn := "host=localhost user=root password=secret dbname=todolistwebapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user := models.Task{Titile: "task1"}

	db.First(&user)

	fmt.Println(user)

}
