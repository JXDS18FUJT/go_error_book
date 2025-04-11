package apiValid

import (
	"347613781qq.com/goInit1/app/http/control/apiControl"
	"github.com/gin-gonic/gin"
)

func QuestionValid(c *gin.Context) {
	apiControl.Question(c)

}
