package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Task, error)
	FindByID(ID uint) (Task, error)
	Create(task Task) (Task, error)
	CreateSubTask(task Task) (Task, error)
	FilterByTitleAndDescription(title, description string, page, limit int) ([]Task, error)
	//Pagination(pagination *dtos.Pagination) (dtos.Pagination, int)
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

func (r *repository) FindByID(ID uint) (Task, error) {
	var parentTaskWithChildren Task
	err := r.db.Preload("Children").First(&parentTaskWithChildren, ID).Error

	return parentTaskWithChildren, err
}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *repository) FilterByTitleAndDescription(title, description string, page, limit int) ([]Task, error) {
	var tasks []Task
	query := r.db.Model(&Task{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	offset := (page - 1) * limit
	if err := query.Preload("Children").Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) CreateSubTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}
