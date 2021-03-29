package main

import (
	"github.com/BurntSushi/toml"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"github.com/zj0395/golib/conf"
	"github.com/zj0395/golib/db"
	"github.com/zj0395/golib/golog"

	"github.com/zj0395/urlshortener/models"
	"github.com/zj0395/urlshortener/router"
)

func initDB() {
	dbConf := conf.DBConf{}
	if _, err := toml.DecodeFile("./config/db.toml", &dbConf); err != nil {
		panic(err)
		return
	}
	if v, err := db.InitDB(&dbConf); err != nil {
		panic(err.Error())
	} else {
		models.InitDB(v)
	}
	golog.Info().Msg("Init db succ.")
}

func initLog() {
	conf := golog.LogConf{
		File:  "log/tmp.log",
		Level: -1,
	}
	golog.SetDefault(golog.Init(&conf))
}

func main() {
	routerObj := routing.New()

	initLog()
	initDB()

	golog.Info().Msg("Init all succ.")

	router.AddRouter(routerObj)
	panic(fasthttp.ListenAndServe(":8082", routerObj.HandleRequest))
}
