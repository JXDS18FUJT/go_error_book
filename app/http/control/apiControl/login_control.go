package apiControl

import (
	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"347613781qq.com/goInit1/app/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var obj = bindtypes.Login{}
	c.ShouldBind(&obj)
	token, _ := service.GenerateToken(obj.Username, obj.Password)

	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
		"data": gin.H{
			"token": token,
		},
	})

}
