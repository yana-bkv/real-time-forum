package models

type Post struct {
	ID     int
	Title  string
	Body   string
	Author string
}

type Comment struct {
	ID     int
	Body   string
	Author string
}
