package task

import "time"

type Service interface {
	FindAll() ([]Task, error)
	FindByID(ID int) (Task, error)
	Create(task Task) (Task, error)
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

func (s *service) FindByID(ID int) (Task, error) {
	task, err := s.repository.FindByID(ID)
	return task, err
}

func (s *service) Create(request AddTaskRequest) (Task, error) {
	datetime := time.Now()
	newTask := Task{
		Title:       request.Title,
		Descryption: request.Descryption,
		CreatedAt:   datetime,
		UpdatedAt:   datetime,
		IsFinished:  false,
		ParentsID:   request.ParentID,
		PriorityID:  request.PriorityID,
	}

	task, err := s.repository.Create(newTask)
	return task, err
}
