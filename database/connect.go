package database

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"jwt-authentication/models"
	"log"
	"strconv"
)

var DB *sql.DB

func ConnectDB() {
	// Initialize db
	var err error
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	DB = db

	// Create user table
	sqlStmtUser := SqlUserDb()
	_, err = db.Exec(sqlStmtUser)
	if err != nil {
		log.Fatal("Error creating table", err, sqlStmtUser)
	}

}

func CreateUser(db *sql.DB, user *models.User) error {
	sqlStmt := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	if user.Username == "" {
		return errors.New("Username is required")
	}

	result, err := db.Exec(sqlStmt, user.Username, user.Email, user.Password)
	if err != nil {
		if isDuplicateEmailError(err) {
			return errors.New("email already taken")
		}
		return errors.New("Error creating user")
	}

	// Adds id from db to json data
	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = uint(userID)

	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	sqlStmt := `SELECT id, username, email, password FROM users WHERE email = ?`
	row := db.QueryRow(sqlStmt, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(db *sql.DB, id string) (*models.User, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Invalid ID format:", err)
	}

	sqlStmt := `SELECT id, username, email, password FROM users WHERE id = ?`
	row := db.QueryRow(sqlStmt, intID)

	var user models.User
	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Check if error is a unique constraint violation for SQLite
func isDuplicateEmailError(err error) bool {
	if err != nil && err.Error() == "UNIQUE constraint failed: users.email" {
		return true
	}
	return false
}
