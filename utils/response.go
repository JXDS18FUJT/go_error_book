package utils

import "github.com/gin-gonic/gin"

func ResError(dataCode int, msg string, data interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})

}
