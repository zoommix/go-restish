package user

var (
	getAllUsersSQL = `SELECT id, first_name, last_name, email, username
FROM users
ORDER BY first_name, last_name
LIMIT $1
OFFSET $2`
)
