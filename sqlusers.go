package sqlusers

var CreateUsersTable = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY,
	name TEXT
)`

var InsertUsers = `
INSERT INTO users (name) VALUES ('Alice'), ('Bob'), ('Charlie'), ('David'), ('Eve')
`

var SelectUsers = `
SELECT id, name FROM users LIMIT 5
`
