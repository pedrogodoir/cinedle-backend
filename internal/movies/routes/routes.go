package routes

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {
	movies := route.Group("/movies")
	{
		movies.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//movies.POST("/", service.CreateMovie) #exemplo
	}
}
