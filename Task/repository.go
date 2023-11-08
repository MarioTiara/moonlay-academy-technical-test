package task

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Task, error)
	FindByID(ID int) (Task, error)
	Create(task Task) (Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *repository) FindByID(ID int) (Task, error) {
	var task Task
	err := r.db.Find(&task).Error
	return task, err
}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}
