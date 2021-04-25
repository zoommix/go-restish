package user

import (
	"context"
	"strings"
)

func (s *Service) GetAllUsers(limit, skip int64) ([]User, error) {
	users := make([]User, 0)

	rows, err := s.DB.Query(context.Background(), getAllUsersSQL, limit, skip)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}

		if err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.Username,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	rows.Close()

	return users, nil
}

func (s *Service) CreateUser(u *User) (*User, error) {
	err := s.DB.QueryRow(
		context.Background(),
		createUserSQL,
		u.FirstName,
		u.LastName,
		strings.ToLower(u.Email),
		u.Username,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}
