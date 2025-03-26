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
	r.GET("nong/:id", apiValid.NongValid)
	api := r.Group("api")
	api.GET("nong/:id", apiValid.NongValid)
	api.GET("login", apiValid.LoginValid)

	api.GET("dylogin", apiValid.DyloginValid)
	api.GET("client_token", apiValid.ClientTokenValid)
	api.GET("schemav2", apiValid.Schemav2Valid)
	api.POST("upload", apiValid.UploadValid)

}
