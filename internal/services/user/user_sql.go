package user

var (
	getAllUsersSQL = `SELECT id, first_name, last_name, email, username, created_at, updated_at
FROM users
ORDER BY first_name, last_name, id DESC
LIMIT $1
OFFSET $2`
)
