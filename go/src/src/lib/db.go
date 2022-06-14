package lib

import "github.com/jinzhu/gorm"

type City struct {
	ID          uint   `gorm:"column:ID" "primary_key"`
	Name        string `gorm:"column:Name"`
	CountryCode string `gorm:"column:CountryCode"`
	District    string `gorm:"column:District"`
	Population  string `gorm:"column:Population"`
}

type Country struct {
}

// SQLConnect DB接続
func Connect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:13306)"
	DBNAME := "world"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
