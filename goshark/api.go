package goshark

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusZHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetHexHandler(ctx *gin.Context) {
	hexValue := ctx.Param("hex")

	p := DecodePacketXML(hexValue, false)
	ctx.JSON(http.StatusOK, p)
}

func GetHexsHandler(ctx *gin.Context) {
	hexValue := ctx.Param("hex")

	p := DecodePacketXML(hexValue, true)
	ctx.JSON(http.StatusOK, p)
}

func HttpServer() {
	r := gin.Default()
	r.GET("/statusz", StatusZHandler)

	r.GET("/api/v1/hex/:hex", GetHexHandler)
	r.GET("/api/v1/hexs/:hex", GetHexsHandler)

	err := r.Run()
	if err != nil {
		log.Fatalf("cannot start http server %s", err)
	}
}
