package models

import (
	"fmt"
	"hash/fnv"
)

const (
	gUrlShortenTablePrefix = "url_shorten_"
	UrlShortenPerTableCnt  = 1 * 2000 * 10000

	// max value: (utils/shorten.maxNum) / UrlShortenPerTableCnt = 1981
	UrlShortenTableCnt = 2

	// for scale up
	urlShortenTableWBegin = 0
)

type UrlShorten struct {
	ID    int64  `gorm:"column:id" json:"id"`
	Url   string `gorm:"column:url" json:"url"`
	Ip    string `gorm:"column:ip" json:"ip"`
	Ctime int64  `gorm:"column:ctime" json:"ctime"`
	Extra string `gorm:"column:extra" json:"extra"`
}

// GetUrlShortenTable for write
func GetUrlShortenTable(url string) string {
	f := fnv.New32()
	f.Write([]byte(url))
	return fmt.Sprintf("%s%d", gUrlShortenTablePrefix, f.Sum32()%(UrlShortenTableCnt-urlShortenTableWBegin)+urlShortenTableWBegin)
}

// GetUrlShortenTableById for read
func GetUrlShortenTableById(id int64) string {
	return fmt.Sprintf("%s%d", gUrlShortenTablePrefix, getUrlShortenTableIdxById(id))
}

func getUrlShortenTableIdxById(id int64) int64 {
	res := id / UrlShortenPerTableCnt
	if res < 0 {
		res = 0
	}
	return res
}
