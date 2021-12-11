package restapi

import "github.com/gin-gonic/gin"

func CreateHealthApi() {

	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"health": "OK",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	router.Run(":50001")
}
