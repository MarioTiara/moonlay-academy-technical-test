package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Task, error)
	FindByID(ID uint) []Task
	Create(task Task) (Task, error)
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

func (r *repository) FindByID(ID uint) []Task {
	var tasks []Task
	//getTaskWithChildren(r.db, nil, &tasks)
	var task Task
	r.db.First(&task, 1)
	tasks = append(tasks, task)
	return tasks

}

func (r *repository) Create(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func getTaskWithChildren(db *gorm.DB, parentID *uint, tasks *[]Task) {
	var children []Task

	query := db.Where("parent_task_id = ?", parentID).Find(&children)
	*tasks = append(*tasks, children...)

	for _, child := range children {
		getTaskWithChildren(query, &child.ID, tasks)
	}
}

// func (r *repository) Pagination(pagination *dtos.Pagination) (interface{}, int) {
// 	// var tasks Task

// 	// totalRows, totalPages, fromRow, toRow := 0, 0, 0, 0

// 	offset := pagination.Page * pagination.Limit

// 	//get data with limit, offest &order
// 	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

// 	//generte where query
// 	searchs := pagination.Searchs

// 	if searchs != nil {
// 		for _, value := range searchs {
// 			column := value.Column
// 			action := value.Action
// 			query := value.Query

// 			switch action {
// 			case "equals":
// 				whereQuery := fmt.Sprintf("%s = ?", column)
// 				find = find.Where(whereQuery, query)
// 				break
// 			case "contains":
// 				whereQuery := fmt.Sprintf("%s LIKE ?", column)
// 				find = find.Where(whereQuery, "%"+query+"%")
// 				break
// 			case "in":

// 			}

// 		}
// 	}

// }
