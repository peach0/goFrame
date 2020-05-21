package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStaffList(c *gin.Context) {
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})

}
