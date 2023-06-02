/**
 * @Author: yy
 * @Description:
 * @File:  timeout
 * @Version: 1.0.0
 * @Date: 2023/06/02 13:46
 */

package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler {
	// 使用函数回调
	return func(c *Context) error {
		log.Println("TimeoutHandler is running...")
		// 这个channel负责通知结束
		finish := make(chan struct{}, 1)
		// 这个channel负责通知panic异常
		panicChan := make(chan any, 1)

		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationCtx)

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 这里做具体的业务
			_ = fun(c)

			// 业务结束后通知结束
			finish <- struct{}{}
		}()
		// 使用 select 关键字来监听三个事件：异常事件、结束事件、超时事件。
		// 请求监听的时候增加锁机制
		select {
		// 监听panic
		case p := <-panicChan:
			c.WriterMux().Lock()
			defer c.WriterMux().Unlock()
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			// 监听结束
			fmt.Println("finish")
		case <-durationCtx.Done():
			// 监听超时
			c.WriterMux().Lock()
			defer c.WriterMux().Unlock()
			c.SetHasTimeout()
			_, _ = c.responseWriter.Write([]byte("timeout"))
		}
		return nil

	}
}
