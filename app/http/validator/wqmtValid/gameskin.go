package wqmtvalid

import (
	wqmtcontrol "347613781qq.com/goInit1/app/http/control/wqmtControl"
	"github.com/gin-gonic/gin"
)

func GameskinValid(c *gin.Context) {
	wqmtcontrol.Gameskin(c)

}
