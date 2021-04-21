package models

import (
	"gorm.io/gorm"
)

var defaultDb *gorm.DB

func GetDB() *gorm.DB {
	return defaultDb
}

func DBTx(table string) *gorm.DB {
	return GetDB().Table(table)
}

func SetDefaultDB(db *gorm.DB) {
	defaultDb = db
}
