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
	appHttp "github.com/Joker-desire/go-web/app/http"
	"github.com/Joker-desire/go-web/app/provider/demo"
	"github.com/Joker-desire/go-web/framework/gin"
	"github.com/Joker-desire/go-web/framework/middleware"
	"github.com/Joker-desire/go-web/framework/provider/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()
	// 绑定具体的服务
	_ = core.Bind(&app.HadeAppProvider{})
	_ = core.Bind(&demo.ServiceProviderDemo{})
	core.Use(
		gin.Recovery(),
		middleware.CostMiddleware())
	appHttp.Routes(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}
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
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用Server.Shutdown()方法来优雅的关闭服务
	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
