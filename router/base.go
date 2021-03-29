package router

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type BaseResponse struct {
	Errno int         `json:"errno"`
	Data  interface{} `json:"data"`
}

func SetOutput(ctx *fasthttp.RequestCtx, data *BaseResponse) {
	body, _ := json.Marshal(data)
	ctx.Response.Header.Set("Content-Type", "application/Json; charset=utf-8")
	ctx.SetBody(body)
}
