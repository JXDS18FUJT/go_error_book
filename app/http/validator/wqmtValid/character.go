package wqmtvalid

import (
	wqmtcontrol "347613781qq.com/goInit1/app/http/control/wqmtControl"
	"github.com/gin-gonic/gin"
)

func CharacterPropertyValid(c *gin.Context) {
	wqmtcontrol.CharacterProperty(c)

}

func CharacterSkillValid(c *gin.Context) {
	wqmtcontrol.CharacterSkill(c)

}
