package routers

import (
	"347613781qq.com/goInit1/app/http/middleware"
	wqmtvalid "347613781qq.com/goInit1/app/http/validator/wqmtValid"
	"github.com/gin-gonic/gin"
)

func InitWqmtRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	wqmt := r.Group("wqmt")
	//角色技能
	wqmt.GET("character_skill", wqmtvalid.CharacterSkillValid)
	//角色属性
	wqmt.GET("character_property", wqmtvalid.CharacterPropertyValid)
	//烙印
	wqmt.GET("brand", wqmtvalid.BrandValid)
	//皮肤
	wqmt.GET("gameskin", wqmtvalid.GameskinValid)
	//成就
	wqmt.GET("achieve", wqmtvalid.AchieveValid)
	//卡池
	wqmt.GET("uppool", wqmtvalid.UppoolValid)
}
