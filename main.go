/**
 * @Author: yy
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/06/01 10:26
 */

package main

import (
	"context"
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		//middleware.TimeoutMiddleware(time.Second),
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
	// 启动服务goroutine
	go func() {
		_ = server.ListenAndServe()
	}()

	// 阻塞等待退出信号
	// 当前的goroutine阻塞等待退出信号
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 阻塞等待退出信号
	<-quit

	// 控制优雅关闭等待的最长时间
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	// 调用Server.Shutdown()方法来优雅的关闭服务
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
