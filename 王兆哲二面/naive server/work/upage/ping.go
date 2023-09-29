package upage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"msg": "pong",
	})
}
