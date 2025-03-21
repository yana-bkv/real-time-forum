package database

func SqlUserDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		sqlAction = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	case "createUser":
		sqlAction = `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	case "getUserByEmail":
		sqlAction = `SELECT id, username, email, password FROM users WHERE email = ? `
	case "getUserById":
		sqlAction = `SELECT id, username, email, password FROM users WHERE id = ?`
	case "getAllUsers":
		sqlAction = `SELECT id, username, email, password FROM users`
	}

	return sqlAction
}

func SqlPostDb() string {
	// Add post table to db
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		user TEXT NOT NULL
	);`
	return sqlStmt
}
