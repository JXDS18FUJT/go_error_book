package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func UserinfoValid(c *gin.Context) {
	apiControl.Userinfo(c)

}
