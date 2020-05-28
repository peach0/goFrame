package goTest

import (
	"gof/bootstrap/conf"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	confList := conf.LoadDatabaseConf()
	t.Logf("%v", confList)
	for _ ,conf := range confList {
		t.Logf("%v", conf)

	}
}
