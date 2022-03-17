package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	blogRouter := r.Group("/blog")
	{
		blogRouter.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "welcome to blog index"})
		})
	}
}
