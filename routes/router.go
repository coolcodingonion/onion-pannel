package routes

import (
	"onion-pannel/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routr := r.Group("/api/v1") 
	{
		routr.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	
	return r
}