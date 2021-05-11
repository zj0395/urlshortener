package main

import (
	"os"
	"strings"
	"text/template"

	"github.com/zj0395/urlshortener/models"
)

// $BQ means backquota'`'
var (
	// url shorten table
	_shortenTableTpl = `
CREATE TABLE IF NOT EXISTS $BQurl_shorten_{{.tableIdx}}$BQ (
  $BQid$BQ bigint(20) NOT NULL AUTO_INCREMENT,
  $BQurl$BQ varchar(256) NOT NULL DEFAULT '' COMMENT 'url',
  $BQip$BQ varchar(41) NOT NULL DEFAULT '' COMMENT 'creater ip',
  $BQctime$BQ int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'create time',
  $BQextra$BQ varchar(256) NOT NULL DEFAULT '' COMMENT 'extra info',
  PRIMARY KEY ($BQid$BQ)
) ENGINE=InnoDB AUTO_INCREMENT={{.autoIncre}} DEFAULT CHARSET=utf8;
`
	shortenTableTpl = strings.ReplaceAll(_shortenTableTpl, "$BQ", "`")

	// access history table
	_accessHistoryTableTpl = `
CREATE TABLE IF NOT EXISTS $BQaccess_history_{{.tableIdx}}$BQ (
  $BQid$BQ bigint(20) NOT NULL AUTO_INCREMENT,
  $BQcode$BQ char(6) NOT NULL COMMENT 'short code',
  $BQip$BQ varchar(41) NOT NULL DEFAULT '' COMMENT 'access user ip',
  $BQctime$BQ int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'access time',
  PRIMARY KEY ($BQid$BQ)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
`
	accessHistoryTableTpl = strings.ReplaceAll(_accessHistoryTableTpl, "$BQ", "`")
)

func main() {
	GenerateShortenTable()
	GenerateAccessHistoryTable()
}

func GenerateShortenTable() {
	tpl, err := template.New("shorten").Parse(shortenTableTpl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < models.UrlShortenTableCnt; i++ {
		data := map[string]interface{}{
			"tableIdx":  i,
			"autoIncre": models.UrlShortenPerTableCnt * i,
		}
		tpl.Execute(os.Stdout, data)
	}
}

func GenerateAccessHistoryTable() {
	tpl, err := template.New("access").Parse(accessHistoryTableTpl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < models.AccessHistoryTableCnt; i++ {
		data := map[string]interface{}{
			"tableIdx": i,
		}
		tpl.Execute(os.Stdout, data)
	}
}
