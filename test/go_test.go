package goTest

import (
	"gof/bootstrap/conf"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	a := conf.LoadDatabaseConf()
	t.Logf("%v", a)
}
