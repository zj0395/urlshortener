package router

import (
	"github.com/fasthttp/router"

	"github.com/zj0395/golib/fhlib"

	"github.com/zj0395/urlshortener/handlers"
)

func GetRouter() *router.Router {
	r := router.New()
	r.ANY("/surl/create", fhlib.NewWrapper(handlers.Create).Exec)
	r.ANY("/{code}", fhlib.NewWrapper(handlers.Access).Exec)
	return r
}
