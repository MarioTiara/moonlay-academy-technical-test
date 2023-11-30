package task

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Descryption string `json:"descryption"`
	ParentID    int    `json:"parents_task_id"`
}
