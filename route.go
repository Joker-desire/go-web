/**
 * @Author: yy
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2023/06/01 14:05
 */

package main

import (
	"github.com/Joker-desire/go-web/framework/gin"
	"github.com/Joker-desire/go-web/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	core.GET("/foo", FooControllerHandler)
	// 在core中使用middleware.Test3() 为单个路由增加中间件
	core.GET("/hello", middleware.Test3(), HelloControllerHandler)
	api := core.Group("/api")
	{
		// 对UserControllerHandler进行超时处理
		api.GET("/user", UserControllerHandler)
		api.GET("/user/:id/detail", UserControllerHandler2)
		v1 := api.Group("/v1")
		{
			v1.GET("/user", UserControllerHandler)
		}
	}
	resp := core.Group("/resp")
	{
		resp.GET("/jsonp", TestJsonP)
		resp.GET("/text", TestText)
		resp.GET("/xml", TestXml)
		resp.GET("/html", TestHtml)
	}
}
