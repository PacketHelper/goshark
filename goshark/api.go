package goshark

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func StatusZHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetHexHandler(ctx *gin.Context) {
	hexValue := ctx.Param("hex")

	hexArray, err := Hex2Array(hexValue)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": OddParity,
		})
		return
	}

	hexdump := DumpHex(hexValue)
	DecodePacket(hexValue)
	ctx.JSON(http.StatusOK, gin.H{
		"hex":     strings.Join(hexArray, " "),
		"hexdump": hexdump,
	})
}

func HttpServer() {
	r := gin.Default()
	r.GET("/statusz", StatusZHandler)

	r.GET("/api/v1/hex/:hex", GetHexHandler)
	err := r.Run()
	if err != nil {
		log.Fatalf("cannot start http server %s", err)
	}
}
