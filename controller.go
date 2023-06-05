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
	"github.com/Joker-desire/go-web/framework"
	"log"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {

	// 获取请求参数
	all := ctx.QueryAll()
	fmt.Printf("%v\n", all)
	device, _ := ctx.QueryString("device", "pc")
	fmt.Printf("device = %v\n", device)
	age, _ := ctx.QueryInt("age", 0)
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
	return nil
}

func HelloControllerHandler(ctx *framework.Context) error {
	ctx.SetOkStatus().Json("hello world")
	return nil
}

func UserControllerHandler(ctx *framework.Context) error {
	time.Sleep(10 * time.Second)
	if name, ok := ctx.QueryString("name", "Joker"); ok {
		ctx.SetOkStatus().Json("hello " + name + "!")
	}
	return nil
}

func UserControllerHandler2(ctx *framework.Context) error {
	param := ctx.Param("id")
	log.Println(param)
	ctx.SetOkStatus().Json("hello world")
	return nil
}

func TestJsonP(ctx *framework.Context) error {
	ctx.SetOkStatus().JsonP("hello JsonP")
	return nil
}

func TestText(ctx *framework.Context) error {
	ctx.SetOkStatus().Text("hello Text")
	return nil
}

func TestXml(ctx *framework.Context) error {
	ctx.SetOkStatus().Xml("hello Xml")
	return nil
}

func TestHtml(ctx *framework.Context) error {
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
	ctx.SetOkStatus().Html("index.html", data)
	return nil
}
