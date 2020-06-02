package goTest

import (
	"github.com/jinzhu/gorm"
	"gof/bootstrap/db"
	"log"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	db.ConnectDb()
	DB := db.GetDB("gof")
	log.Print()
	DB.Create(&Product{Code: "L1212", Price: 1000})
}


type Product struct {
	gorm.Model
	Code string
	Price uint
}
