package priority

import task "github.com/marioTiara/todolistwebapi/Task"

type Service interface {
	Create(request AddPriorotyTaskRequest) (task.Priority, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(request AddPriorotyTaskRequest) (task.Priority, error) {
	newPriority := task.Priority{
		Priority: request.PriorityName,
	}

	priority, err := s.repository.Create(newPriority)
	return priority, err
}
