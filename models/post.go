package models

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorId int    `json:"author_id"`
	Time     string `json:"time"`
}

type CreatePostRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	AuthorID    int    `json:"author_id"`
	CategoryIDs []int  `json:"category_ids"`
}
