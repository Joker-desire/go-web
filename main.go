/**
 * @Author: yy
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/06/01 10:26
 */

package main

import (
	"github.com/Joker-desire/go-web/app/console"
	appHttp "github.com/Joker-desire/go-web/app/http"
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/provider/app"
	"github.com/Joker-desire/go-web/framework/provider/kernel"
	"log"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()
	// 绑定具体的服务
	_ = container.Bind(&app.HadeAppProvider{})

	// 将HTTP引擎初始化，并且作为服务提供者绑定到服务容器中
	if engine, err := appHttp.NewHttpEngine(); err == nil {
		_ = container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	if err := console.RunCommand(container); err != nil {
		log.Fatal(err)
	}
}
