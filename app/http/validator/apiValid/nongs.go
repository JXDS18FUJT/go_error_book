package apiValid

import (
	bindtypes "347613781qq.com/goInit1/app/http/bindTypes"
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
)

func NongValid(c *gin.Context) {

	var obj = bindtypes.Nongs{}
	zhcn.RegisterGlobal()
	bindError := c.ShouldBind(&obj)
	if bindError == nil {

		v := validate.Struct(obj)
		zhcn.Register(v)
		//通过检测
		if v.Validate() {
			apiControl.Nong(c)

		} else {
			utils.ResError(402, v.Errors.One(), gin.H{
				"id": obj.Id,
			}, c)

		}

	}

}
