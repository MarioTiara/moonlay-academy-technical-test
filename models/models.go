package models

import (
	"time"
)

type Task struct {
	ID          int       `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Descryption string    `gorm:"column:descryption"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	IsFinished  bool      `gorm:"column:is_finished"`
	ParentsID   int       `gorm:"column:parents_id"`

	PriorityID int `gorm:"column:priority_id"`

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
