package wqmtvalid

import (
	wqmtcontrol "347613781qq.com/goInit1/app/http/control/wqmtControl"
	"github.com/gin-gonic/gin"
)

func UppoolValid(c *gin.Context) {
	wqmtcontrol.Uppool(c)

}
