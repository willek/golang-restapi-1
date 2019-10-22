package model

type Blog struct {
	Id    int
	Title string
	Text  string
}

type BlogResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Blog
}
