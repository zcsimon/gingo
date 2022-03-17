package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	blogRouter := r.Group("/shop")
	{
		blogRouter.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "welcome to shop index"})
		})
	}
}
