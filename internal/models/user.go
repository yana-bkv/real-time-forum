package models

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Login
}

type Login struct {
	SessionToken string
	CSRFToken    string
}
