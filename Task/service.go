package task

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
	//datetime := time.Now()

	var parentTask = convertRequestToTaskEntity(request)
	for _, task := range request.Children {
		parentTask.Children = append(parentTask.Children, convertRequestToTaskEntity(task))
	}

	task, err := s.repository.Create(parentTask)
	return task, err
}

func convertRequestToTaskEntity(request AddTaskRequest) Task {
	newtask := Task{Title: request.Title, Descryption: request.Descryption}
	return newtask
}
