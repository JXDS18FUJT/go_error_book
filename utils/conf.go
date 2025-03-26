package utils

import "gopkg.in/ini.v1"

type AppConf struct {
	Port      int    `ini:"port"`
	Release   bool   `ini:"release"`
	JwtSecret string `ini:"jwtSecret"`
}
type MysqlConf struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Db       string `ini:"db"`
}
type RedisConf struct {
	user     string `ini:"user"`
	Password string `ini:"password"`
	Addr     string `ini:"addr"`
	Db       string `ini:"db"`
}

var AppConfig = AppConf{}
var MysqlConfig = MysqlConf{}
var RedisConfig = RedisConf{}

func InitAppConf() error {
	return ini.MapTo(&AppConfig, "conf/app.ini")
}
func InitMysqlConf() error {
	return ini.MapTo(&MysqlConfig, "conf/mysql.ini")
}

func InitRedisConf() error {
	return ini.MapTo(&RedisConfig, "conf/redis.ini")
}
