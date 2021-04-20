package user

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	DB *pgxpool.Pool
}

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Username  string
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{
		DB: db,
	}
}

type StudentService interface {
	GetAllUsers(limit, skip uint) ([]User, error)
	GetUserByID(ID uint) (User, error)
	GetUserByUsername(username string) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(ID uint, newUser User) (User, error)
	DeleteUser(ID uint) error
}

func (s *Service) GetAllUsers(limit, skip uint) ([]User, error) {
	var users []User

	rows, err := s.DB.Query(context.Background(), getAllUsersSQL, limit, skip)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}

		if err := rows.Scan(u.ID, u.FirstName, u.LastName, u.Username); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	rows.Close()

	return users, nil
}
