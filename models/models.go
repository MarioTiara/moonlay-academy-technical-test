package models

import "gorm.io/gorm"

type Task struct {
	*gorm.Model
	Titile      string
	Descryption string
	IsFinished  bool

	PriorityID int
	Priority   Priority

	Files []File
}

type Priority struct {
	*gorm.Model
	Priority_name string

	Tasks []Task
}

type File struct {
	*gorm.Model
	Extension string
	Filename  string

	TaskID int
	Task   Task
}

// type Users struct {
// 	Id       int
// 	Name     string
// 	Username string
// }

// type User struct {
// 	gorm.Model
// 	Name      string
// 	CompanyID int
// 	Company   Company
// }

// type Company struct {
// 	ID        int
// 	Name      string
// 	CountryID int
// 	Country   Country
// }

// type Country struct {
// 	ID   int
// 	Name string
// }
