package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/zj0395/golib/fhlib"
	"github.com/zj0395/golib/liberr"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils/shorten"
)

func Create(ctx *fasthttp.RequestCtx) {
	logger := fhlib.GetLogger(ctx)
	url := strings.TrimSpace(string(ctx.QueryArgs().Peek("url")))

	if url == "" {
		logger.Error().Str("error", "empty url").Msg("Create shorturl error")
		fhlib.SetError(ctx, liberr.ParamError)
		return
	}

	// insert db
	obj := models.UrlShorten{
		Url:   url,
		Ip:    fhlib.ClientIP(ctx),
		Ctime: time.Now().Unix(),
	}
	dbRes := models.DBTx(models.GetUrlShortenTable(url)).Create(&obj)
	if dbRes.Error != nil {
		logger.Error().Str("error", dbRes.Error.Error()).Msg("Create shorturl error")
		fhlib.SetError(ctx, dbRes.Error)
		return
	}
	if dbRes.RowsAffected == 0 {
		logger.Error().Msg("Create shorturl fail, affectet row = 0")
		fhlib.SetError(ctx, liberr.DBError)
		return
	}

	// idshorten
	code := shorten.IDShorten(obj.ID)
	logger.Info().Int64("id", obj.ID).Str("code", code).Msg("Create url succ")

	host := ctx.Request.Header.Host()
	resp := map[string]interface{}{
		"code": code,
		"surl": fmt.Sprintf("%s/%s", host, code),
	}
	fhlib.SetData(ctx, resp)
}
