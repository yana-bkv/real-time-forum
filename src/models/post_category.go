package models

type PostCategory struct {
	PostID      int   `json:"post_id"`
	CategoryIDs []int `json:"category_ids"`
}
