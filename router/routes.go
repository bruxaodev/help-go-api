package router

import (
	"help/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	scream := handler.Scream{}
	v1 := router.Group("/api/v1")
	{
		v1.GET("/scream", scream.Get)
		v1.POST("/scream", scream.Post)
		v1.DELETE("/scream", scream.Delete)
		v1.PUT("/scream", scream.Put)
		v1.GET("/screamers", scream.GetList)
	}
}
