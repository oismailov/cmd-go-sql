package models

import (
	"cmd-go-sql/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error
var DBSession *gorm.DB

//InitDatabaseSession - creates connection to mysql database
func InitDatabaseSession() {
	DBSession, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.Cfg.DatabaseSettings.DatabaseUsername,
		config.Cfg.DatabaseSettings.DatabasePassword,
		config.Cfg.DatabaseSettings.DatabaseHost,
		config.Cfg.DatabaseSettings.DatabasePort,
		config.Cfg.DatabaseSettings.DatabaseName,
		"utf8",
		"True",
		"Local",
	))

	if err != nil {
		panic(err)
	}
}

//GetDatabaseSession returns a databse session to use
func GetDatabaseSession() *gorm.DB {
	if DBSession.DB().Ping() != nil {
		InitDatabaseSession()
	}
	return DBSession
}
