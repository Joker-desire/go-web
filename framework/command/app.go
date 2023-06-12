/**
 * @Author: yy
 * @Description:
 * @File:  app
 * @Version: 1.0.0
 * @Date: 2023/06/09 18:58
 */

package command

import (
	"context"
	"github.com/Joker-desire/go-web/framework/cobra"
	"github.com/Joker-desire/go-web/framework/contract"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initAppCommand() *cobra.Command {
	appCommand.AddCommand(appStartCommand)
	return appCommand
}

var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  `业务应用控制命令，其包含业务启动，关闭，重启，查询等命令`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 打印帮助文档
		_ = cmd.Help()
		return nil
	},
}

var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := cmd.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		// 创建一个server服务
		server := &http.Server{
			Handler: core,
			Addr:    ":8080",
		}
		// 启动server服务
		go func() {
			_ = server.ListenAndServe()
		}()

		// 当前的goroutine等待信号量
		quit := make(chan os.Signal)
		// 监控信号：SIGINT, SIGTERM, SIGQUIT
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		// 阻塞等待退出信号
		<-quit

		// 调用Server.Shutdown graceful关闭服务
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		return nil
	},
}
