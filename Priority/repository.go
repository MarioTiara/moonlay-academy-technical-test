package priority

import (
	task "github.com/marioTiara/todolistwebapi/Task"
	"gorm.io/gorm"
)

type Repository interface {
	Create(priority task.Priority) (task.Priority, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(priority task.Priority) (task.Priority, error) {
	err := r.db.Create(&priority).Error

	return priority, err
}
