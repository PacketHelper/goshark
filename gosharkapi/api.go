package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpServer() {
	r := gin.Default()
	r.GET(
		"/status", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	r.Run()
}
