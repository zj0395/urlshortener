package models

import (
	"fmt"
	"hash/fnv"
)

const (
	gUrlShortenTablePrefix = "url_shorten_"
	UrlShortenIdBegin      = 50 * 10000 * 10000
	UrlShortenPerTableCnt  = 1 * 2000 * 10000

	// max value: 2590
	// (utils/shorten.maxNum - UrlShortenIdBegin) / UrlShortenPerTableCnt
	UrlShortenTableCnt = 2
)

type UrlShorten struct {
	ID    int64  `gorm:"column:id" json:"id"`
	Url   string `gorm:"column:url" json:"url"`
	Ip    string `gorm:"column:ip" json:"ip"`
	Ctime int64  `gorm:"column:ctime" json:"ctime"`
	Extra string `gorm:"column:extra" json:"extra"`
}

func GetUrlShortenTable(url string) string {
	f := fnv.New32()
	f.Write([]byte(url))
	return fmt.Sprintf("%s%d", gUrlShortenTablePrefix, f.Sum32()%UrlShortenTableCnt)
}

func GetUrlShortenTableById(id int64) string {
	return fmt.Sprintf("%s%d", gUrlShortenTablePrefix, getUrlShortenTableIdxById(id))
}

func getUrlShortenTableIdxById(id int64) int64 {
	res := (id - UrlShortenIdBegin) / UrlShortenPerTableCnt
	if res < 0 {
		res = 0
	}
	return res
}
