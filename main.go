package main

import (
	"os"
	"shProxy/conf"
	"shProxy/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	//获取端口号 因为端口号有时不固定
	port := os.Getenv("PORT")
	if port == "" {
		// 没有获取到端口号 则使用3000端口
		port = "3000"
	}
	r.Run(":" + port)
}
