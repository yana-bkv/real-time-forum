package models

type Post struct {
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Time       string `json:"time"`
	AuthorName string `json:"author_name"`
	//Likes    uint      `json:"likes"`
	//Comments []Comment `json:"comments"`
}

type Like struct {
	Id         uint   `json:"id"`
	AuthorName string `json:"author_name"`
}

type Comment struct {
	Id       uint   `json:"id"`
	AuthorId uint   `json:"author_id"`
	Body     string `json:"body"`
	Time     string `json:"time"`
}
