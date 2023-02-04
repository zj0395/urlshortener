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
	}
	if v, err := db.InitDB(&dbConf); err != nil {
		panic(err.Error())
	} else {
		models.SetDefaultDB(v)
	}
	golog.Info().Msg("Init db succ.")
}

func initLog() {
	conf := &golog.RotateLogConf{
		LogConf: golog.LogConf{
			File:  "log/tmp.log",
			Level: -1,
		},

		EnableRotate: true,
		Period:       24,
		UseMinute:    false,
	}
	golog.SetDefault(golog.NewRotateLog(conf).GetLogger())
}

func main() {
	initLog()
	initDB()

	golog.Info().Msg("Init all succ.")

	server := &fasthttp.Server{
		Name:             "url-shorten",
		Handler:          router.GetRouter().Handler,
		ReadTimeout:      time.Second * 3,
		WriteTimeout:     time.Second * 5,
		DisableKeepalive: false,
		LogAllErrors:     true,
		Logger:           golog.LogForwarder(),
	}

	go func() {
		err := server.ListenAndServe(fmt.Sprintf(":%d", 8083))
		if err != nil {
			panic(err)
		}
	}()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	// Stop the service gracefully.
	golog.Info().Msg("begin shutdown")

	// may cost sometime
	// depends on `DisableKeepalive` and `ReadTimeout` options
	server.Shutdown()

	golog.Info().Msg("shutdown succ")
}
