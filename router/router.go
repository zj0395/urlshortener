package router

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func AddRouter(router *routing.Router) {
	router.Get("/test", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

	router.Any("/create", Create)
	router.Any("/<code>", Access)
}
