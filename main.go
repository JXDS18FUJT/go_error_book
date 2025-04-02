package main

import (
	"fmt"

	//必须引入到项目里347613781qq.com/goInit1/docs初始化

	_ "347613781qq.com/goInit1/docs"
	"347613781qq.com/goInit1/routers"
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           go_my_init
// @version         1.0
// @description     This is a simple go template.

// @contact.name   API Support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9091
// @BasePath /api/

// @title           go_my_init
// @version         1.0
// @description     This is a simple go template.

// @contact.name   API Support

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9091
// @BasePath /api/

func main() {

	//先执行获取配置
	utils.InitMysqlConf()
	utils.InitAppConf()

	utils.InitRedisConf()
	// utils.InitLogger()
	DbError := utils.InitMySQL(utils.MysqlConfig.User, utils.MysqlConfig.Password, utils.MysqlConfig.Host, utils.MysqlConfig.Port, utils.MysqlConfig.Db)
	if DbError != nil {
		fmt.Print("数据库连接错误")
		return
	}

	RedisError := utils.InitRedis(utils.RedisConfig.Addr, utils.RedisConfig.Password, 0)
	if RedisError != nil {
		fmt.Print("redis连接错误")
		return
	}
	// utils.Redis.Set(utils.RedisCtx, "wxname2", "张三", 0)
	var r = gin.Default()
	routers.InitApiRouter(r)
	// var user = model.User{
	// 	Username:     "admin",
	// 	PasswordHash: "123456",
	// 	Status:       1,
	// }
	// model.CreateUser(&user)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//格式化
	r.Run(fmt.Sprintf(":%d", utils.AppConfig.Port))

}
