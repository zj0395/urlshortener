module github.com/zj0395/urlshortener

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/fasthttp/router v1.3.11
	github.com/valyala/fasthttp v1.44.0
	github.com/zj0395/golib v0.0.0-20230204174546-faf4ef21b1bd
	gorm.io/gorm v1.21.9
)

// replace github.com/zj0395/golib => ../golib
