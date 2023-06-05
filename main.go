/**
 * @Author: yy
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/06/01 10:26
 */

package main

import (
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	core := framework.NewCore()
	// 使用use注册中间件
	//core.Use(
	//	middleware.Test1(),
	//	middleware.Test2(),
	//	middleware.TimeoutMiddleware(time.Second*10))
	core.Use(
		middleware.RecoveryMiddleware(),
		middleware.CostMiddleware(),
		middleware.TimeoutMiddleware(time.Second),
	)
	// 注册路由
	registerRouter(core)
	// 打印路由表
	routers := core.GetRouters()
	log.Printf("%T\n", routers)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core,
		// 请求监听地址
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
