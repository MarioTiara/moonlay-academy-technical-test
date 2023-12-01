package task

import "time"

type Service interface {
	FindAll() ([]Task, error)
	FindByID(ID uint) []Task
	Create(request AddTaskRequest) (Task, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Task, error) {
	tasks, err := s.repository.FindAll()
	return tasks, err
}

func (s *service) FindByID(ID uint) []Task {
	task := s.repository.FindByID(ID)
	return task
}

func (s *service) Create(request AddTaskRequest) (Task, error) {
	datetime := time.Now()
	var newTask = Task{}
	if request.ParentID <= 0 {

		newTask = Task{
			Title:       request.Title,
			Descryption: request.Descryption,
			CreatedAt:   datetime,
		}

	} else {
		var parentID = uint(request.ParentID)
		newTask = Task{
			Title:       request.Title,
			Descryption: request.Descryption,
			CreatedAt:   datetime,
			ParentID:    &parentID,
		}
	}

	task, err := s.repository.Create(newTask)
	return task, err
}
