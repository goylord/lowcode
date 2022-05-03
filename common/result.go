package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResultError(code int, message string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}

func ResultSuccess(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func ResultSuccessList(data interface{}, total int64, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"total": total,
		"data":  data,
	})
}
