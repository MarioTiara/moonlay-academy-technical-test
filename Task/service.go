package task

type Service interface {
	FindAll() ([]Task, error)
	FindByID(ID uint) (Task, error)
	CreateSubTask(parentID uint, request AddTaskRequest) (Task, error)
	Create(request AddTaskRequest) (Task, error)
	FilterTask(title, description string, page, limit int) ([]Task, error)
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

func (s *service) FindByID(ID uint) (Task, error) {
	task, err := s.repository.FindByID(ID)
	return task, err
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

func (s *service) CreateSubTask(parentID uint, request AddTaskRequest) (Task, error) {
	var subTask = Task{Title: request.Title, Descryption: request.Descryption, ParentID: &parentID}
	return s.repository.CreateSubTask(subTask)
}

func (s *service) FilterTask(title, description string, page, limit int) ([]Task, error) {
	return s.repository.FilterByTitleAndDescription(title, description, page, limit)
}

func convertRequestToTaskEntity(request AddTaskRequest) Task {
	newtask := Task{Title: request.Title, Descryption: request.Descryption}
	return newtask
}
