package apiValid

import (
	"bytes"
	"io"

	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"347613781qq.com/goInit1/app/http/control/apiControl"

	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

func LoginValid(c *gin.Context) {
	var obj = bindtypes.Login{}
	zhcn.RegisterGlobal()
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(200, gin.H{"error": "读取请求失败"})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	bindError := c.ShouldBind(&obj)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if bindError == nil {
		v := validate.Struct(obj)
		zhcn.Register(v)
		//通过检测
		if v.Validate() {
			apiControl.Login(c)

		} else {
			utils.ResError(402, v.Errors.One(), nil, c)

		}

	}

}
