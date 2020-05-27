package db

import (
	"gof/bootstrap/conf"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

const (
	DRIVER_MY_SQL  = "mysql"
	DRIVER_SQLITE3 = "sqlite3"
)

func ConnectDb() {
	confList := conf.LoadDatabaseConf()
	for conf := range confList {
		switch conf.DbType {
		case DRIVER_MY_SQL:
			db = connectDbMysql(
				conf.Url,
				conf.Port,
				conf.Uname,
				conf.Passwd,
				conf.DbName)
		default:
			db = connectDbMysql(
				conf.Url,
				conf.Port,
				conf.Uname,
				conf.Passwd,
				conf.DbName)
		}

	}
}

func connectDbMysql(url, port, uname, passwd, dbName string) *gorm.DB {

}
