package models

type Comment struct {
	Id       int    `json:"id"`
	Body     string `json:"body"`
	PostId   int    `json:"post_id"`
	AuthorId int    `json:"author_id"`
	Time     string `json:"time"`
}
