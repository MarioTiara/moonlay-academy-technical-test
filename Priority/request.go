package priority

type AddPriorotyTaskRequest struct {
	PriorityName string `json:"priority_name" binding:"required"`
}
