package goshark

import (
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

func CreateRouter() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/statusz", StatusZHandler)

	router.GET("/api/v1/hex/:hex", GetHexHandler)
	return
}
