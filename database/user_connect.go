package database

import (
	"database/sql"
	"errors"
	"jwt-authentication/models"
	"log"
	"strconv"
)

func CreateUser(db *sql.DB, user *models.User) error {
	sqlStmt := SqlUserDb("createUser")

	if user.Username == "" {
		return errors.New("username is required")
	}

	result, err := db.Exec(sqlStmt, user.Username, user.Email, user.Password)
	if err != nil {
		if isDuplicateEmailError(err) {
			return errors.New("email already taken")
		}
		return errors.New("error creating user")
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
	sqlStmt := SqlUserDb("getUserByEmail")
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

	sqlStmt := SqlUserDb("getUserById")
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
