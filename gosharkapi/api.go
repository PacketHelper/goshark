package gosharkapi

import (
	"log"
	"net/http"

	"github.com/PacketHelper/goshark/v2/goshark"
	"github.com/gin-gonic/gin"
)

func statusZ(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getHex(ctx *gin.Context) {
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
}

func HttpServer() {
	r := gin.Default()
	r.GET("/statusz", statusZ)

	r.GET("/api/v1/hex/:hex", getHex)
	err := r.Run()
	if err != nil {
		log.Fatalf("cannot start http server %s", err)
	}
}
