package task

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Descryption string `json:"descryption"`
	ParentID    int    `json:"parent_id"`
	PriorityID  int    `json:"priority_id"`
}