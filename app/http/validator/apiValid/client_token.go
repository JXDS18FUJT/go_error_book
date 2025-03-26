package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func ClientTokenValid(c *gin.Context) {
	apiControl.ClientToken(c)

}
