package goshark

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	return router
}

func TestStatusZHandler(t *testing.T) {
	mockResponse := `{"message":"ok"}`
	r := SetUpRouter()
	r.GET("/statusz", StatusZHandler)

	req, _ := http.NewRequest("GET", "/statusz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

func TestGetHexHandler(t *testing.T) {
	mockResponse := `{"hex":"00 11 aa","hexdump":["0000 00 11 aa"]}`

	r := SetUpRouter()
	r.GET("/api/v1/hex/:hex", GetHexHandler)

	req, _ := http.NewRequest("GET", "/api/v1/hex/0011aa", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
