package repositories

import (
	"database/sql"
	"errors"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"log"
	"strconv"
)

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	sqlStmt := database.SqlUserDb("createUser")

	if user.Username == "" {
		return errors.New("username is required")
	}

	result, err := database.DB.Exec(sqlStmt, user.Username, user.Email, user.Password)
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
	sqlStmt := database.SqlUserDb("getUserByEmail")
	row := db.QueryRow(sqlStmt, email)

	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	sqlStmt := database.SqlUserDb("getUserByUsername")
	row := db.QueryRow(sqlStmt, username)

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

	sqlStmt := database.SqlUserDb("getUserById")
	row := db.QueryRow(sqlStmt, intID)

	var user models.User
	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	sqlStmt := database.SqlUserDb("getAllUsers")
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Check if error is a unique constraint violation for SQLite
func isDuplicateEmailError(err error) bool {
	if err != nil && err.Error() == "UNIQUE constraint failed: users.email" {
		return true
	}
	return false
}
