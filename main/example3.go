package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由参数
func main() {
	router := gin.Default()

	// 这个handler 将会匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是, 这个将匹配 /user/john/ 以及 /user/john/send
	// 如果没有其他路由器匹配 /user/john, 它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// 对于每个匹配的请求，上下文将保留路由定义
	router.POST("/user/:name/*action", func(c *gin.Context) {
		//c.FullPath() == "/user/:name/*action" // true
	})

	router.Run(":8080")
}
