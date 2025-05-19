package models

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
	AuthorId int    `json:"author_id"`
	Time     string `json:"time"`
}
