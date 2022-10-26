package middleware

import (
	"github.com/valyala/fasthttp"
	"github.com/zj0395/golib/golog"

	"github.com/zj0395/urlshortener/utils"
)

func Init(ctx *fasthttp.RequestCtx) {
	const requestIDHeader = "X-Request-Id"
	var logid string
	if v := ctx.Request.Header.Peek(requestIDHeader); v != nil {
		logid = string(v)
	} else {
		logid = utils.GenLogId()
	}
	utils.SetLogId(ctx, logid)
	logger := golog.GetDefault().With().Str("logid", logid).Logger()
	utils.SetLogger(ctx, &logger)
}
