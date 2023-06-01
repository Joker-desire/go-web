/**
 * @Author: yy
 * @Description:
 * @File:  controller
 * @Version: 1.0.0
 * @Date: 2023/06/01 11:52
 */

package main

import (
	"context"
	"fmt"
	"github.com/Joker-desire/go-web/framework"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {

	// 获取请求参数
	all := ctx.QueryAll()
	fmt.Printf("%v\n", all)
	device := ctx.QueryString("device", "pc")
	fmt.Printf("device = %v\n", device)
	age := ctx.QueryInt("age", 0)
	fmt.Printf("age = %v\n", age)

	// 1. 生成一个超时的Context
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	defer cancel()

	//2. 创建一个新的Goroutine来处理业务逻辑
	// 这个channel负责通知结束
	finish := make(chan struct{}, 1)
	// 这个channel负责通知panic异常
	panicChan := make(chan any, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// 这里做具体的业务
		//time.Sleep(10 * time.Second)

		_ = ctx.Json(200, "hello world")
		// 业务结束后通知结束
		finish <- struct{}{}
	}()
	// 使用 select 关键字来监听三个事件：异常事件、结束事件、超时事件。
	// 请求监听的时候增加锁机制
	select {
	// 监听panic
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		_ = ctx.Json(500, p)
	case <-finish:
		// 监听结束
		fmt.Println("finish")
	case <-durationCtx.Done():
		// 监听超时
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		_ = ctx.Json(500, "time out")
		ctx.SetHasTimeout()

	}
	return nil
}

func HelloControllerHandler(ctx *framework.Context) error {
	return ctx.Json(200, "hello world")
}

func UserControllerHandler(ctx *framework.Context) error {
	name := ctx.QueryString("name", "Joker")
	return ctx.Json(200, "hello "+name+"!")
}
