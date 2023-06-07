/**
 * @Author: yy
 * @Description:
 * @File:  timeout
 * @Version: 1.0.0
 * @Date: 2023/06/02 15:17
 */

package middleware

import (
	"context"
	"github.com/Joker-desire/go-web/framework/gin"
	"log"
	"time"
)

func TimeoutMiddleware(d time.Duration) gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		log.Println("TimeoutMiddleware is running...")
		// 这个channel负责通知结束
		finish := make(chan struct{}, 1)
		// 这个channel负责通知panic异常
		panicChan := make(chan any, 1)

		// 执行业务逻辑前预操作：初始化超时context
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 使用Next执行具体的业务逻辑
			c.Next()

			// 业务结束后通知结束
			finish <- struct{}{}
		}()
		// 使用 select 关键字来监听三个事件：异常事件、结束事件、超时事件。
		// 请求监听的时候增加锁机制
		select {
		// 监听panic
		case p := <-panicChan:
			c.ISetStatus(500).IJson("timeout")
			log.Println(p)
		case <-finish:
			// 监听结束
			log.Println("timeout finish")
			break
		case <-durationCtx.Done():
			// 监听超时
			c.ISetStatus(500).IJson("timeout")
		}
	}
}
