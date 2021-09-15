package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"` //omitempty nil or default
	Msg  string      `json:"msg,omitempty"`
}


const (
	SUCCESS  = 2000
	ERROR    = 4000
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code: ERROR,
		Msg:  "404 not found",
	})
}

func CommonFailed(message string, code int, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Msg:  message,
	})
}

func CommonSuccess(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  message,
	})
}
