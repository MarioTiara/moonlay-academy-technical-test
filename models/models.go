package models

import (
	"time"
)

type Task struct {
	ID          int
	Title       string
	Descryption string
	CreatedAt   time.Time
	UpdatdAt    time.Time
	IsFinished  bool
	ParentsID   int

	PriorityID int
	Priority   Priority

	Files []Files
}

type Priority struct {
	ID       int
	Priority string
	Tasks    []Task
}

type Files struct {
	ID        int
	Extension string
	Filename  string

	TaskID int
}
