/**
 * @Author: yy
 * @Description:
 * @File:  cost
 * @Version: 1.0.0
 * @Date: 2023/06/02 15:27
 */

package middleware

import (
	"github.com/Joker-desire/go-web/framework/gin"
	"log"
	"time"
)

func CostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		startTime := time.Now()
		// 执行具体的业务逻辑
		c.Next()
		// 记录结束时间
		endTime := time.Now()
		// 计算耗时
		costTime := endTime.Sub(startTime)
		// 打印耗时
		path := c.Request.RequestURI
		method := c.Request.Method
		log.Printf("path: %s, method: %s, cost: %v", path, method, costTime)
	}
}
