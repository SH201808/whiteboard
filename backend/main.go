package main

import (
	"github.com/gin-contrib/pprof"
	"log"
	"whiteboard/dao/mysql"
	"whiteboard/dao/redis"
	"whiteboard/middleware/rabbitmq"
	"whiteboard/router"
	"whiteboard/setting"
	"whiteboard/utils/validator"
)

func init() {
	err := setting.Init()
	if err != nil {
		log.Panicln("配置文件错误:", err)
	}
	//初始化MySQL数据库
	mysql.Init(setting.Conf.MySQLConfig)
	//初始化Reids数据库
	redis.Init(setting.Conf.RedisConfig)
	//初始化Validate
	validator.Init()
	rabbitmq.Init()
}

func main() {
	// 注册路由
	r := router.SetupRouter()
	//性能测试
	pprof.Register(r)

	r.Run(":8080")
}
