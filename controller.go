/**
 * @Author: yy
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2023/06/01 11:52
 */

package main

import (
	"fmt"
	"github.com/Joker-desire/go-web/framework/gin"
	"github.com/Joker-desire/go-web/provider/demo"
	"log"
	"time"
)

func FooControllerHandler(ctx *gin.Context) {

	// 获取请求参数
	all := ctx.QueryAll()
	fmt.Printf("%v\n", all)
	device, _ := ctx.DefaultQueryString("device", "pc")

	fmt.Printf("device = %v\n", device)
	age, _ := ctx.DefaultQueryInt("age", 0)
	fmt.Printf("age = %v\n", age)
	time.Sleep(10 * time.Second)
	//
	//// 1. 生成一个超时的Context
	//durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	//defer cancel()
	//
	////2. 创建一个新的Goroutine来处理业务逻辑
	//// 这个channel负责通知结束
	//finish := make(chan struct{}, 1)
	//// 这个channel负责通知panic异常
	//panicChan := make(chan any, 1)
	//go func() {
	//	defer func() {
	//		if p := recover(); p != nil {
	//			panicChan <- p
	//		}
	//	}()
	//	// 这里做具体的业务
	//	//time.Sleep(10 * time.Second)
	//
	//	_ = ctx.Json(200, "hello world")
	//	// 业务结束后通知结束
	//	finish <- struct{}{}
	//}()
	//// 使用 select 关键字来监听三个事件：异常事件、结束事件、超时事件。
	//// 请求监听的时候增加锁机制
	//select {
	//// 监听panic
	//case p := <-panicChan:
	//	ctx.WriterMux().Lock()
	//	defer ctx.WriterMux().Unlock()
	//	_ = ctx.Json(500, p)
	//case <-finish:
	//	// 监听结束
	//	fmt.Println("finish")
	//case <-durationCtx.Done():
	//	// 监听超时
	//	ctx.WriterMux().Lock()
	//	defer ctx.WriterMux().Unlock()
	//	_ = ctx.Json(500, "time out")
	//	ctx.SetHasTimeout()
	//
	//}
}

func HelloControllerHandler(ctx *gin.Context) {
	ctx.ISetOkStatus().IJson("hello world")
}

func UserControllerHandler(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	if name, ok := ctx.DefaultQueryString("name", "Joker"); ok {
		ctx.ISetOkStatus().IJson("hello " + name + "!")
	}
}

func UserControllerHandler2(ctx *gin.Context) {
	param := ctx.Param("id")
	log.Println(param)
	ctx.ISetOkStatus().IJson("hello world")
}

func TestJsonP(ctx *gin.Context) {
	ctx.ISetOkStatus().IJsonP("hello JsonP")
}

func TestText(ctx *gin.Context) {
	ctx.ISetOkStatus().IText("hello Text")
}

func TestXml(ctx *gin.Context) {
	ctx.ISetOkStatus().IXml("hello Xml")
}

func TestHtml(ctx *gin.Context) {
	type Todo struct {
		Title string
		Done  bool
	}
	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	ctx.ISetOkStatus().IHtml("index.html", data)
}

func TestDemoService(c *gin.Context) {
	// 获取demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)
	// 调用服务实例的方法
	foo := demoService.GetFoo()
	// 返回结果
	c.ISetOkStatus().IJson(foo)
}
