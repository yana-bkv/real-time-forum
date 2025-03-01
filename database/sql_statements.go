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

func SqlPostDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		// Add post table to db
		sqlAction = `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT UNIQUE NOT NULL,
		category TEXT NOT NULL,
		time TEXT NOT NULL,
		author_name TEXT NOT NULL,
		likes INTEGER,
		comments TEXT
);`
	case "createPost":
		sqlAction = `INSERT INTO posts (title, body, time, author_name) VALUES (?, ?, ?, ?)`
	case "getPostById":
		sqlAction = `SELECT id, title, body, author_name FROM posts WHERE id = ?`
	case "getAllPosts":
		sqlAction = `SELECT id, title, body, author_name FROM posts`
	case "getAllPostsByCategory":
		sqlAction = `SELECT id, title, body, author_name FROM posts WHERE category = ?`
	case "deletePost":
		sqlAction = `DELETE FROM posts WHERE id = ?`
	}
	return sqlAction
}
