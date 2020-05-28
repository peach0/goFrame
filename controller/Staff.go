package controller

import (
	"gof/action/staff"

	"github.com/gin-gonic/gin"
)

func Staff(E *gin.Engine) {
	g := E.Group("staff")
	{
		g.GET("/getstafflist", staff.GetStaffList)
	}
}

func handler(mapFunc map[string]string) {

}
