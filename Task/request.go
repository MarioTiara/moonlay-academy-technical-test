package task

type AddTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Descryption string `json:"description"`
	FileName    string `json:"file_name"`
	Children    []AddTaskRequest
}

type AddSublistByID struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
