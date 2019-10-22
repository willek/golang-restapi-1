package model

type Blog struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

type BlogResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []Blog `json:"data"`
}
