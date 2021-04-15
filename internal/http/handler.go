package http

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

// Handler ...
type Handler struct {
	Router *router.Router
}

// NewHandler ...
func NewHandler() *Handler {
	return &Handler{}
}

// InitRouter ...
func (h *Handler) InitRouter() *Handler {
	h.Router = router.New()

	h.Router.GET("/api/status", Status)

	return h
}

// Status ...
func Status(c *fasthttp.RequestCtx) {
	WriteJSON(c, map[string]string{"status": "ok"})
	SetStatus(c, fasthttp.StatusOK)
}
