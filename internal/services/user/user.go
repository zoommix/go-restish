package user

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Service struct {
	DB *pgxpool.Pool
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserJSON struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (u *User) ToJSON() UserJSON {
	return UserJSON(*u)
}
