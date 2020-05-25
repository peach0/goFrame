package goTest

import (
"gof/library/util/db"
"testing"
)

func TestHelloWorld(t *testing.T) {
	a := db.LoadDatabaseConf()
	t.Logf("%v", a)
}
