module github.com/zj0395/urlshortener

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/qiangxue/fasthttp-routing v0.0.0-20160225050629-6ccdc2a18d87
	github.com/valyala/fasthttp v1.22.0
	github.com/zj0395/golib v0.0.0-20200619060749-84d5e8891798
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)

replace github.com/zj0395/golib => ../golib
