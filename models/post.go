package models

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Category string `json:"category"`
	AuthorId int    `json:"author_id"`
	Time     string `json:"time"`
}

type Comment struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	PostId   int    `json:"post_id"`
	AuthorId int    `json:"author_id"`
	Time     string `json:"time"`
}

type Like struct {
	Id       int `json:"id"`
	PostId   int `json:"post_id"`
	AuthorId int `json:"author_id"`
}
