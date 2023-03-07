package handlers

import (
	"time"

	"github.com/valyala/fasthttp"

	"github.com/zj0395/golib/fhlib"
	"github.com/zj0395/golib/liberr"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils/errors"
	"github.com/zj0395/urlshortener/utils/shorten"
)

type AccessReq struct {
	Code string `json:"form:code"`
}

func Access(ctx *fasthttp.RequestCtx) {
	logger := fhlib.GetLogger(ctx)
	code, ok := ctx.UserValue("code").(string)
	if !ok || len(code) != 6 {
		logger.Warn().Str("code", code).Msg("Invalid code")
		fhlib.SetErrorOutput(ctx, liberr.ParamError)
		return
	}

	// recover id from code
	id := shorten.IDRecover(code)

	if !models.IsIDValid(id) {
		logger.Warn().Str("code", code).Msg("invalid code")
		fhlib.SetErrorOutput(ctx, errors.CodeNotExist)
		return
	}

	// get from db
	obj := models.UrlShorten{}
	models.DBTx(models.GetUrlShortenTableById(id)).Find(&obj, id)
	if obj.ID == 0 {
		logger.Warn().Str("code", code).Int64("id", id).Msg("Code not found")
		fhlib.SetErrorOutput(ctx, errors.CodeNotExist)
		return
	}

	clientIP := fhlib.ClientIP(ctx)
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

	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Redirect(obj.Url, fasthttp.StatusFound)
}
