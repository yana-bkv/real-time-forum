package models

type Like struct {
	Id       int `json:"id"`
	PostId   int `json:"post_id"`
	AuthorId int `json:"author_id"`
}
