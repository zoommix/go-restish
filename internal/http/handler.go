package http

import (
	"go-restish/internal/services/user"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

type Handler struct {
	Router      *router.Router
	UserService *user.Service
}

func NewHandler(userService *user.Service) *Handler {
	return &Handler{UserService: userService}
}

// InitRouter ...
func (h *Handler) InitRouter() *Handler {
	h.Router = router.New()

	h.Router.GET("/api/status", h.Status)
	h.Router.GET("/api/users", h.ListUsers)
	h.Router.POST("/api/users", h.CreateUser)

	return h
}

func (h *Handler) Status(c *fasthttp.RequestCtx) {
	WriteJSON(c, map[string]string{"status": "ok"})
	SetStatus(c, fasthttp.StatusOK)
}

func (h *Handler) ListUsers(c *fasthttp.RequestCtx) {
	users, err := h.UserService.GetAllUsers(
		Paging(c, "limit", 15),
		Paging(c, "skip", 0),
	)

	if err != nil {
		log.Println(err)
		SetStatus(c, fasthttp.StatusInternalServerError)

		return
	}

	jsonUsers := make([]user.UserJSON, 0)

	for _, u := range users {
		jsonUsers = append(jsonUsers, u.ToJSON())
	}

	WriteJSON(c, jsonUsers)
	SetStatus(c, fasthttp.StatusOK)
}

func (h *Handler) CreateUser(c *fasthttp.RequestCtx) {
	u, err := h.UserService.ParseUserFromCtx(c)

	if err != nil {
		WriteError(c, fasthttp.StatusUnprocessableEntity, newRespError(err.Error()))
		return
	}

	u, err = h.UserService.CreateUser(u)

	if err != nil {
		WriteError(c, fasthttp.StatusInternalServerError, newRespError(err.Error()))
		return
	}

	WriteJSON(c, u.ToJSON())
}
