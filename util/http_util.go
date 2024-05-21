package util

import (
	"github.com/gin-gonic/gin"
)

func ErrorJson(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"message": err.Error(),
	})
}
