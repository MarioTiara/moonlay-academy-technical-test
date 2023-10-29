package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Pets []Pet
}

type Pet struct {
	gorm.Model
	Name   string
	User   User
	UserID uint
}

type Employee struct {
	Id   int
	Name string
}
