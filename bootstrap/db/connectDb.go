package db

import (
	"fmt"
	"gof/bootstrap/conf"
	"log"

	"github.com/jinzhu/gorm"
)

var DBMap map[string]*gorm.DB

const (
	DRIVER_MY_SQL  = "mysql"
	DRIVER_SQLITE3 = "sqlite3"
)

func GetDB(dbName string) *gorm.DB {
	if db, ok := DBMap[dbName]; !ok {
		return nil
	} else {
		return db
	}
}

func ConnectDb() {
	confList := conf.LoadDatabaseConf()
	for _, conf := range confList {
		var DB *gorm.DB
		switch conf.DbType {
		case DRIVER_MY_SQL:
			DB = connectDbMysql(
				conf.Url,
				conf.Port,
				conf.Uname,
				conf.Passwd,
				conf.DbName,
				conf.CharSet)
		default:
			DB = connectDbMysql(
				conf.Url,
				conf.Port,
				conf.Uname,
				conf.Passwd,
				conf.DbName,
				conf.CharSet)
		}
		DBMap[conf.DbName] = DB

	}
}

func connectDbMysql(url, port, uname, passwd, dbName, charset string) *gorm.DB {

	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		uname,
		passwd,
		url,
		port,
		dbName,
		charset,
	)

	db, err := gorm.Open(DRIVER_MY_SQL, dns)
	if err != nil {
		log.Fatalf("models.InitDbMySQL err: %v", err)
	}
	db.SingularTable(true)
	return db
}
