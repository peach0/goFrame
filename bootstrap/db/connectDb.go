package db

import (
	"sync"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var lock *sync.Mutex = &sync.Mutex{}

const (
	MYSQL = iota
	SQLSERVER
)

//使用多例模式来生成数据路链接
//什么时候关闭
