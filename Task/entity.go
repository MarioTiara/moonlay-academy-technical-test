package task

import "time"

type Task struct {
	ID          int       `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Descryption string    `gorm:"column:descryption"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	ParentsID   int       `gorm:"column:parents_task_id"`
}

type Files struct {
	ID       int    `gorm:"column:file_id"`
	FileName string `gorm:"file_name"`
	TaskID   int    `gorm:"task_id"`
}
