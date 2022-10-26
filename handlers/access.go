package handlers

import (
	"time"

	"github.com/valyala/fasthttp"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils"
	"github.com/zj0395/urlshortener/utils/errors"
	"github.com/zj0395/urlshortener/utils/shorten"
)

type AccessReq struct {
	Code string `json:"form:code"`
}

func Access(ctx *fasthttp.RequestCtx) {
	logger := utils.GetLogger(ctx)
	code, ok := ctx.UserValue("code").(string)
	if !ok || len(code) != 6 {
		logger.Warn().Str("code", code).Msg("Invalid code")
		SetErrorOutput(ctx, errors.ParamError)
		return
	}

	// recover id from code
	id := shorten.IDRecover(code)

	// get from db
	obj := models.UrlShorten{}
	models.DBTx(models.GetUrlShortenTableById(id)).Find(&obj, id)
	if obj.ID == 0 {
		logger.Warn().Str("code", code).Int64("id", id).Msg("Code not found")
		SetErrorOutput(ctx, errors.CodeNotExist)
		return
	}

	clientIP := utils.ClientIP(ctx)
	go func() {
		ah := models.AccessHistory{
			Code:  code,
			Ip:    clientIP,
			Ctime: time.Now().Unix(),
		}
		dbRes := models.DBTx(models.GetAccessHistoryTableById(id)).Create(&ah)
		if dbRes.Error != nil {
			logger.Warn().Str("error", dbRes.Error.Error()).Msg("Create access history error")
		}
		if dbRes.RowsAffected == 0 {
			logger.Warn().Msg("Create access history fail, affectet row = 0")
		}
	}()

	logger.Info().Int64("id", id).Str("code", code).Msg("access succ")
	ctx.Redirect(obj.Url, fasthttp.StatusFound)
}
