package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func CollyValid(c *gin.Context) {
	apiControl.Colly(c)

}
