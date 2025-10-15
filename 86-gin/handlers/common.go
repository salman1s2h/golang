package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
}
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pong")
}
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{"health": "ok"})
}
