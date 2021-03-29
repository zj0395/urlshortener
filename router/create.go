package router

import (
	"time"

	routing "github.com/qiangxue/fasthttp-routing"

	"github.com/zj0395/golib/golog"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/utils"
	"github.com/zj0395/urlshortener/utils/shorten"
)

type CreateResp struct {
	Code string
}

func Create(c *routing.Context) error {
	url := string(c.QueryArgs().Peek("url"))

	obj := models.UrlShorten{
		Url:   url,
		CTime: time.Now().Unix(),
		IP:    utils.ClientIP(c.RequestCtx),
	}
	dbRes := models.DBTx(models.GetUrlShortenTable(url)).Create(&obj)
	if dbRes.Error != nil {
		golog.Error().Str("error", dbRes.Error.Error()).Msg("Create shorturl error")
		return dbRes.Error
	}
	if dbRes.RowsAffected == 0 {
		golog.Error().Msg("Create shorturl fail, affectet row = 0")
		return dbRes.Error
	}
	code := shorten.IDShorten(obj.ID)
	golog.Info().Int64("id", obj.ID).Str("code", code).Msg("Create url succ")

	resp := BaseResponse{
		Data: CreateResp{
			Code: code,
		},
	}
	SetOutput(c.RequestCtx, &resp)
	return nil
}
