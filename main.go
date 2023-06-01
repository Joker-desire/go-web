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
	"net/http"
)

func main() {
	core := framework.NewCore()
	// 注册路由
	registerRouter(core)
	core.Get("foo", FooControllerHandler)
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
