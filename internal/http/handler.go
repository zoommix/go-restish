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

	return h
}

func (h *Handler) Status(c *fasthttp.RequestCtx) {
	users, err := h.UserService.GetAllUsers(15, 0)

	if err != nil {
		log.Println(err)
	}

	log.Println(users)

	WriteJSON(c, map[string]string{"status": "ok"})
	SetStatus(c, fasthttp.StatusOK)
}
