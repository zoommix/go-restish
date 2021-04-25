package user

import (
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
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

type UserService interface {
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

func (s *Service) ParseUserFromCtx(c *fasthttp.RequestCtx) (*User, error) {
	params := UserJSON{}
	err := json.Unmarshal(c.PostBody(), &params)

	if err != nil {
		return nil, err
	}

	u := &User{
		Email:     params.Email,
		Username:  params.Username,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	return u, nil
}
