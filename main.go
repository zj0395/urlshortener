package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
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
		models.SetDefaultDB(v)
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
	initLog()
	initDB()

	golog.Info().Msg("Init all succ.")

	server := &fasthttp.Server{
		Name:             "url-shorten",
		Handler:          router.GetRouter().Handler,
		ReadTimeout:      time.Second * 30,
		WriteTimeout:     time.Second * 30,
		DisableKeepalive: false,
		LogAllErrors:     true,
		Logger:           golog.GetDefault(),
	}

	go func() {
		err := server.ListenAndServe(fmt.Sprintf(":%d", 8082))
		if err != nil {
			panic(err)
			return
		}
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// Stop the service gracefully.
	golog.Info().Msg("begin shutdown")
	server.Shutdown()
	golog.Info().Msg("shutdown succ")
}
