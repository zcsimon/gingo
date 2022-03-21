package blog

import (
	"box/middle"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	blogRouter := r.Group("/blog", middle.Middleware())
	{
		blogRouter.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": "200", "message": "welcome to blog index"})
		})
	}

}
