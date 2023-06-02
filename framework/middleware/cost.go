/**
 * @Author: yy
 * @Description:
 * @File:  cost
 * @Version: 1.0.0
 * @Date: 2023/06/02 15:27
 */

package middleware

import (
	"github.com/Joker-desire/go-web/framework"
	"log"
	"time"
)

func CostMiddleware() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 记录开始时间
		startTime := time.Now()
		// 执行具体的业务逻辑
		_ = c.Next()
		// 记录结束时间
		endTime := time.Now()
		// 计算耗时
		costTime := endTime.Sub(startTime)
		// 打印耗时
		path := c.GetRequest().RequestURI
		method := c.GetRequest().Method
		log.Printf("path: %s, method: %s, cost: %v", path, method, costTime)
		return nil
	}
}
