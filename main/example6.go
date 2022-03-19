package main

import "github.com/gin-gonic/gin"

func main() {
	//记录日志的颜色
	gin.ForceConsoleColor()

	// 使用默认中间件创建一个 gin路由:
	// logger 与 recovery (crash-free) 中间件
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
