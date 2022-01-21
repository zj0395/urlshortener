package handlers

import (
	"strings"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils"
	"github.com/zj0395/urlshortener/utils/errors"
	"github.com/zj0395/urlshortener/utils/shorten"
)

func Create(ctx *fasthttp.RequestCtx) {
	logger := utils.GetLogger(ctx)
	url := strings.TrimSpace(string(ctx.QueryArgs().Peek("url")))

	if url == "" {
		logger.Error().Str("error", "empty url").Msg("Create shorturl error")
		utils.SetError(ctx, errors.ParamError)
		return
	}

	// insert db
	obj := models.UrlShorten{
		Url:   url,
		Ip:    utils.ClientIP(ctx),
		Ctime: time.Now().Unix(),
	}
	dbRes := models.DBTx(models.GetUrlShortenTable(url)).Create(&obj)
	if dbRes.Error != nil {
		logger.Error().Str("error", dbRes.Error.Error()).Msg("Create shorturl error")
		utils.SetError(ctx, dbRes.Error)
		return
	}
	if dbRes.RowsAffected == 0 {
		logger.Error().Msg("Create shorturl fail, affectet row = 0")
		utils.SetError(ctx, errors.DbError)
		return
	}

	// idshorten
	code := shorten.IDShorten(obj.ID)
	logger.Info().Int64("id", obj.ID).Str("code", code).Msg("Create url succ")

	resp := map[string]interface{}{
		"code": code,
	}
	utils.SetData(ctx, resp)
}
