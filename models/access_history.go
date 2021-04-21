package models

import (
	"fmt"
)

const (
	gAccessHistoryTablePrefix = "access_history_"
	gAccessHistoryTableCnt    = 1
)

type AccessHistory struct {
	ID    int64  `gorm:"column:id" json:"id"`
	Code  string `gorm:"column:code" json:"code"`
	Ip    string `gorm:"column:ip" json:"ip"`
	Ctime int64  `gorm:"column:ctime" json:"ctime"`
}

func GetAccessHistoryTableById(id int64) string {
	return fmt.Sprintf("%s%d", gAccessHistoryTablePrefix, getUrlShortenTableIdxById(id)%gAccessHistoryTableCnt)
}
