package reply

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	response := Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

func Error(c *gin.Context, code int, message string) {
	response := Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(code, response)
}
