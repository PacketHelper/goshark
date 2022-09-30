package gosharkapi

import (
	"net/http"

	"github.com/PacketHelper/goshark/v2/goshark"
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

	r.GET(
		"/api/v1/hex/:hex", func(ctx *gin.Context) {
			hexValue := ctx.Param("hex")

			hexArray, err := goshark.Hex2Array(hexValue)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "wrong packet",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"hex": hexArray,
			})
		})
	r.Run()
}
