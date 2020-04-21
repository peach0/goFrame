package main

import (
	"fmt"
	"gof/engine"
	"net/http"
)

func main() {
	e := engine.New()
	e.GET("/", indexHander)

	http.ListenAndServe(":8888", e)
}

func indexHander(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "urll.Path = %q\n", req.URL.Path)
}
