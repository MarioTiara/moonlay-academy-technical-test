package models

import (
	"gorm.io/gorm"
)

type task struct {
	gorm.Model
	title       string
	descryption string
	perent      *task
	parent_id   int `gorm:"TYPE:integer REFERENCE task"`
	isFinished  bool
	priorityId  int
	priority    priority
}

type priority struct {
	gorm.Model
	priority_name string
}

type file struct {
	gorm.Model
	extension string
	filename  string
	taskid    int
}
