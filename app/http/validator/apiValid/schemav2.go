package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func Schemav2Valid(c *gin.Context) {
	apiControl.Schemav2(c)

}
