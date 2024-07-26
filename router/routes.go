package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/scream", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Heeeeeeeelp-me!!!!!!!!!",
			})
		})
		v1.POST("/scream", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Heeeeeeeelp-me!!!!!!!!!",
			})
		})
		v1.DELETE("/scream", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Heeeeeeeelp-me!!!!!!!!!",
			})
		})
		v1.PUT("/scream", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Heeeeeeeelp-me!!!!!!!!!",
			})
		})
		v1.GET("/screamers", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Heeeeeeeelp-me!!!!!!!!!",
			})
		})
	}
}
