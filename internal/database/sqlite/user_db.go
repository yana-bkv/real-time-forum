package sqlite

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"log/slog"
	"real-time-forum/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func InitUserDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	username Text,
    	email Text,
    	password Text
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Error creating table", err, sqlStmt)
	}

}

func CreateUser(username string, email string, password string) {
	if len(username) < 3 || len(email) < 3 || len(password) < 3 {
		slog.Error("Invalid username or email or password")
		return
	}

	// Check if the username already exists
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? OR email = ?", username, email).Scan(&count)
	if err != nil {
		slog.Error("Error checking if user exists:", err)
		return
	}

	if count > 0 {
		slog.Error("Username or email already taken")
		return
	}

	// Hash the password before storing
	hashedPassword, err := hashPassword(password)
	if err != nil {
		slog.Error("Error hashing password:", err)
	}

	_, err = DB.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", username, email, hashedPassword)
	if err != nil {
		slog.Error("Error inserting user into the database:", err)
		return
	}
	slog.Info("User created successfully")
}

func ShowUsers() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, username FROM users")
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			slog.Error(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CheckLogin(username string, password string) (models.User, error) {
	rows, err := DB.Query("SELECT id, username, password FROM users WHERE username = ?", username)
	if err != nil {
		slog.Error("Database query error:", "error", err)
		return models.User{}, err
	}
	defer rows.Close()

	// Check if user exists
	if !rows.Next() {
		slog.Warn("User not found:", "username", username)
		return models.User{}, errors.New("user not found")
	}

	// Scan the result into a variable
	var dbID int
	var dbUsername, dbEmail, dbPassword string
	if err := rows.Scan(&dbID, &dbUsername, &dbEmail, &dbPassword); err != nil {
		slog.Error("Error scanning user:", "error", err)
		return models.User{}, err
	}

	if !checkPasswordHash(password, dbPassword) {
		err = errors.New("invalid password")
		return models.User{}, err
	}

	// Log successful retrieval
	slog.Info("User found:", "username", dbUsername)

	return models.User{
		ID:       dbID,
		Username: dbUsername,
		Email:    dbEmail,
		Password: dbPassword,
	}, nil
}

//func DeleteUser(username string) error {
//
//}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
