package router

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"github.com/zj0395/golib/golog"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils/shorten"
)

func Access(c *routing.Context) error {
	code := c.Param("code")
	if len(code) != 6 {
		golog.Warn().Str("code", code).Msg("Invalid code")
		return fmt.Errorf("param error")
	}

	id := shorten.IDRecover(code)
	obj := models.UrlShorten{
		ID: id,
	}
	models.DBTx(models.GetUrlShortenTableById(id)).Find(&obj, id)
	c.Redirect(obj.Url, fasthttp.StatusFound)
	return nil
}
