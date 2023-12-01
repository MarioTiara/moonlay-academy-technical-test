package task

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Descryption string `json:"descryption"`
	Children    []AddTaskRequest
}
