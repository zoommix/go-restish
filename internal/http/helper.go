package http

import (
	"encoding/json"
	"strconv"

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

func Paging(c *fasthttp.RequestCtx, key string, def int64) int64 {
	d := def
	bytes := c.QueryArgs().Peek(key)

	if len(bytes) == 0 {
		return d
	}

	d, err := strconv.ParseInt(string(bytes), 10, 64)

	if err != nil {
		return d
	}

	return d
}
