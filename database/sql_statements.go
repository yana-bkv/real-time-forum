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
		author_id INTEGER NOT NULL,
		time TEXT NOT NULL
);`
	case "createPost":
		sqlAction = `INSERT INTO posts (title, body, author_id, time) VALUES (?, ?, ?, ?)`
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

func SqlCategoryDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		sqlAction = `
CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    text TEXT UNIQUE NOT NULL
);`
	case "createCategory":
		sqlAction = `INSERT INTO categories (text) VALUES (?)`
	case "getCategoryById":
		sqlAction = `SELECT id, text FROM categories WHERE id = ?`
	case "getAllCategories":
		sqlAction = `SELECT id, text FROM categories`
	case "deleteCategory":
		sqlAction = `DELETE FROM categories WHERE id = ?`
	}
	return sqlAction
}

func SqlPostCategoryDb(action string) string {
	var sqlAction string
	switch action {
	case "createTable":
		sqlAction = `
CREATE TABLE IF NOT EXISTS post_categories (
    post_id INTEGER,
    category_id INTEGER,
 	PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);`
	case "assignCategoryToPost":
		sqlAction = `INSERT INTO post_categories (post_id, category_id) VALUES (?, ?)`
	case "unassignCategoryFromPost":
		sqlAction = `DELETE FROM post_categories WHERE category_id = ? AND post_id = ?`
	case "unassignCategoryFromAllPosts":
		sqlAction = `DELETE FROM post_categories WHERE category_id = ?`
	case "getTagByPostId":
		sqlAction = `SELECT category_id FROM post_categories WHERE post_id = ?`
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
		user_id INTEGER NOT NULL
);`
	case "addLike":
		sqlAction = `INSERT INTO likes (post_id, user_id) VALUES (?, ?)`
	case "getLikesByPostId":
		sqlAction = `SELECT id, post_id, user_id FROM likes WHERE post_id = ?`
	case "deleteLike":
		sqlAction = `DELETE FROM likes WHERE user_id = ? AND post_id = ?;`
	case "getUserCount":
		sqlAction = `SELECT COUNT(user_id) FROM likes WHERE post_id = ?`
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
    created_at TEXT NOT NULL
);`
	case "saveMessage":
		sqlAction = `INSERT INTO messages (sender, receiver, content, created_at) VALUES (?, ?, ?, ?)`
	case "getMessage":
		sqlAction = `SELECT sender, receiver, content, created_at
        FROM messages
        WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)
        ORDER BY created_at ASC`
	}
	return sqlAction
}
