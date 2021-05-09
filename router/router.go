package router

import (
	"runtime/debug"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/zj0395/urlshortener/handlers"
	"github.com/zj0395/urlshortener/handlers/middleware"
	"github.com/zj0395/urlshortener/utils"
	"github.com/zj0395/urlshortener/utils/errors"
)

// control exec of handlers
type wrapper struct {
	middleWares []fasthttp.RequestHandler
	final       fasthttp.RequestHandler
}

func NewWrapper(h fasthttp.RequestHandler) *wrapper {
	return &wrapper{
		middleWares: []fasthttp.RequestHandler{},
		final:       h,
	}
}

func (t *wrapper) Add(handler fasthttp.RequestHandler) *wrapper {
	t.middleWares = append(t.middleWares, handler)
	return t
}

func (t *wrapper) Exec(ctx *fasthttp.RequestCtx) {
	startTime := time.Now()

	middleware.Init(ctx)
	logger := utils.GetLogger(ctx)

	defer func() {
		if p := recover(); p != nil {
			fatalErr := string(debug.Stack())
			logger.Fatal().Str("stack", fatalErr).Msg("[FATAL 500]")
			handlers.SetErrorOutput(ctx, errors.PanicError)
		}
		logger.Info().Int64("costms", time.Since(startTime).Milliseconds()).Msg("Request Done")
	}()

	for _, h := range t.middleWares {
		h(ctx)
		// stop if already error
		if err := utils.GetError(ctx); err != nil {
			handlers.SetErrorOutput(ctx, err)
			return
		}
	}
	if t.final != nil {
		t.final(ctx)
	}
	if err := utils.GetError(ctx); err != nil {
		handlers.SetErrorOutput(ctx, err)
	} else if data := utils.GetData(ctx); data != nil {
		handlers.SetOutput(ctx, data)
	}
}

func GetRouter() *router.Router {
	r := router.New()
	r.ANY("/create", NewWrapper(handlers.Create).Exec)
	r.ANY("/s{code}", NewWrapper(handlers.Access).Exec)
	return r
}
