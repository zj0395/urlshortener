package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"github.com/zj0395/urlshortener/utils"
	"github.com/zj0395/urlshortener/utils/errors"
)

type BaseResponse struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Logid string      `json:"logid"`
	Data  interface{} `json:"data"`
}

type ErrorResponse struct {
	Errno int    `json:"errno"`
	Msg   string `json:"msg"`
	Logid string `json:"logid"`
}

func SetErrorOutput(ctx *fasthttp.RequestCtx, err error) {
	errData := errors.FormatError(err)

	resp := ErrorResponse{
		Errno: errData.Errno,
		Msg:   errData.Msg,
		Logid: utils.GetLogId(ctx),
	}

	body, _ := json.Marshal(resp)
	ctx.Response.Header.Set("Content-Type", "application/Json; charset=utf-8")
	ctx.SetBody(body)
	ctx.SetStatusCode(errData.HttpStatus)
}

func SetOutput(ctx *fasthttp.RequestCtx, data interface{}) {
	resp := BaseResponse{
		Errno: 0,
		Msg:   "succ",
		Logid: utils.GetLogId(ctx),
		Data:  data,
	}

	body, _ := json.Marshal(resp)
	ctx.Response.Header.Set("Content-Type", "application/Json; charset=utf-8")
	ctx.SetBody(body)
}
