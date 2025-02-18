package database

func SqlUserDb() string {
	// Add user table to db
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	return sqlStmt
}
