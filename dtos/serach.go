package dtos

type Serach struct {
	Column string `json:"column"`
	Action string `josn:"action"`
	Query  string `json:"query"`
}
