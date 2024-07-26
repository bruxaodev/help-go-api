package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Scream struct{}

func (s *Scream) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Heeeeeeeelp-me!!!!!!!!!",
	})
}
func (s *Scream) Post(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Heeeeeeeelp-me!!!!!!!!!",
	})
}
func (s *Scream) Put(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Heeeeeeeelp-me!!!!!!!!!",
	})
}
func (s *Scream) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Heeeeeeeelp-me!!!!!!!!!",
	})
}
func (s *Scream) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Heeeeeeeelp-me!!!!!!!!!",
	})
}
