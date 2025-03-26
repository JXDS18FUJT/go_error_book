package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func DyloginValid(c *gin.Context) {
	apiControl.Dylogin(c)

}
