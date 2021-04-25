package user

var (
	getAllUsersSQL = `SELECT id, first_name, last_name, email, username, created_at, updated_at
FROM users
ORDER BY first_name, last_name, id DESC
LIMIT $1
OFFSET $2`

	createUserSQL = `
INSERT INTO users (first_name, last_name, email, username)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at
`
)
