package http

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func SetStatus(ctx *fasthttp.RequestCtx, code int) {
	ctx.Response.SetStatusCode(code)
}

func WriteJSON(ctx *fasthttp.RequestCtx, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		WriteError(ctx, fasthttp.StatusInternalServerError, newRespError(err.Error()))
	}
}

func WriteError(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.SetStatusCode(code)

	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		panic(err)
	}
}

type errorResp map[string]string

func newRespError(msg string) errorResp {
	return errorResp{"error": msg}
}
