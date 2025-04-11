package apiControl

import (
	"fmt"

	"347613781qq.com/goInit1/app/model"
	"github.com/gin-gonic/gin"
)

func Question(c *gin.Context) {
	var questions []model.Question
	err := c.ShouldBindJSON(&questions)
	fmt.Println(questions)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "402",
			"msg":  "参数错误",
		})
		return
	}
	for _, val := range questions {
		model.CreateQuestion(&val)

	}

	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
	})

}
