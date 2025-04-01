package apiControl

import (
	"347613781qq.com/goInit1/app/model"
	"347613781qq.com/goInit1/app/service"
	"github.com/gin-gonic/gin"
)

// @Summary 上传图片的接口
// @Description get string by ID
// @Accept png,jpeg
// @Produce  json
// @Param  file formData file true  "文件"
// @Param name formData  string false "名字"

// @Success 200 {object} model.Nong	"ok"
// @Failure 400 {object} global.AppBusinessError
// @Router /upload [post]
func Userinfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := service.ParseToken(token)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "",
		})
		return

	}
	user, err := model.GetUser(claims.Username)
	c.JSON(200, gin.H{
		"code": 400,
		"msg":  "",
		"data": user,
	})

}
