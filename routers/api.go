package routers

import (
	"347613781qq.com/goInit1/app/http/middleware"
	"347613781qq.com/goInit1/app/http/validator/apiValid"

	"github.com/gin-gonic/gin"
)

func InitApiRouter(r *gin.Engine) {
	//启动
	//router := gin.Default()
	//分析性能
	//pprof.Register(router)
	//启动跨域
	r.Use(middleware.Cors())
	r.POST("login", apiValid.LoginValid)
	api := r.Group("api")
	//开启登录检测
	api.Use(middleware.Token())
	// api.GET("tag/list", apiValid.NongValid)
	api.GET("userInfo", apiValid.UserinfoValid)
	api.GET("dylogin", apiValid.DyloginValid)
	api.GET("client_token", apiValid.ClientTokenValid)
	api.GET("schemav2", apiValid.Schemav2Valid)
	api.POST("upload", apiValid.UploadValid)

}
