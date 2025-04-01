package apiControl

import (
	"context"
	"time"

	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"347613781qq.com/goInit1/app/model"
	"347613781qq.com/goInit1/app/service"
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var obj = bindtypes.Login{}

	c.ShouldBind(&obj)
	user, err := model.GetUser(obj.Username)

	if user.Username != obj.Username {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "用户不存在或密码错误",
		})
		return

	}
	if user.PasswordHash != obj.Password {
		c.JSON(400, gin.H{
			"code": "400",
			"msg":  "用户不存在或密码错误",
		})
		return

	}
	token, _ := service.GenerateToken(obj.Username, obj.Password)
	// 存入 Redis，并设置过期时间
	username := obj.Username
	ctx := context.Background()
	err = utils.Redis.Set(ctx, "token:"+username, token, 2*time.Hour).Err()
	if err != nil {

	}

	c.JSON(200, gin.H{
		"code": "200",
		"msg":  "成功",
		"data": gin.H{
			"token": token,
		},
	})

}
