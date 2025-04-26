package database

func SqlUserDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		sqlAction = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	case "createUser":
		sqlAction = `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	case "getUserByEmail":
		sqlAction = `SELECT id, username, email, password FROM users WHERE email = ? `
	case "getUserByUsername":
		sqlAction = `SELECT id, username, email, password FROM users WHERE username = ? `
	case "getUserById":
		sqlAction = `SELECT id, username, email, password FROM users WHERE id = ?`
	case "getAllUsers":
		sqlAction = `SELECT id, username, email FROM users`
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
		category TEXT,
		author_id INTEGER NOT NULL,
		time TEXT NOT NULL
);`
	case "createPost":
		sqlAction = `INSERT INTO posts (title, body, category, author_id, time) VALUES (?, ?, ?, ?, ?)`
	case "getPostById":
		sqlAction = `SELECT id, title, body, author_id, time FROM posts WHERE id = ?`
	case "getAllPosts":
		sqlAction = `SELECT id, title, body, author_id, time FROM posts`
	case "getAllPostsByCategory":
		sqlAction = `SELECT id, title, body, author_id, time FROM posts WHERE category = ?`
	case "deletePost":
		sqlAction = `DELETE FROM posts WHERE id = ?`
	}
	return sqlAction
}

func SqlCommentDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		// Add post table to db
		sqlAction = `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		body TEXT UNIQUE NOT NULL, -- текст комментария
		post_id INTEGER NOT NULL, -- к какому посту комментарий
		author_id INTEGER NOT NULL, -- кто автора комментария
		time TEXT NOT NULL
);`
	case "createComment":
		sqlAction = `INSERT INTO comments (body, post_id, author_id, time) VALUES (?, ?, ?, ?)`
	case "getCommentsByPostId":
		sqlAction = `SELECT id, body, post_id, author_id, time FROM comments WHERE post_id = ?`
	case "deleteComment":
		sqlAction = `DELETE FROM comments WHERE id = ?`
	}
	return sqlAction
}

func SqlLikeDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		// Add post table to db
		sqlAction = `
	CREATE TABLE IF NOT EXISTS likes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL UNIQUE
);`
	case "addLike":
		sqlAction = `INSERT INTO likes (post_id, author_id) VALUES (?, ?)`
	case "getLikesByPostId":
		sqlAction = `SELECT id, post_id, author_id FROM likes WHERE post_id = ?`
	case "deleteLike":
		sqlAction = `DELETE FROM likes WHERE id = ?`
	}
	return sqlAction
}

func SqlMessageDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		sqlAction = `
	CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sender TEXT NOT NULL,
    receiver TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);`
	case "saveMessage":
		sqlAction = `INSERT INTO likes (sender_name, receiver_name, content) VALUES (?, ?)`
	case "getMessage":
		sqlAction = `SELECT sender, receiver, content, created_at
        FROM messages
        WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)
        ORDER BY created_at ASC`
	}
	return sqlAction
}
