package main

import (
	"gof/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	E := gin.Default()
	controller.InitController(E)
	E.Run() // listen and serve on 0.0.0.0:8080
}
