/**
 * @Author: yy
 * @Description:
 * @File:  recovery
 * @Version: 1.0.0
 * @Date: 2023/06/02 15:23
 */

package middleware

import "github.com/Joker-desire/go-web/framework"

func RecoveryMiddleware() framework.ControllerHandler {
	return func(c *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic异常
		defer func() {
			if err := recover(); err != nil {
				c.WriterMux().Lock()
				defer c.WriterMux().Unlock()
				_ = c.Json(500, err)
			}
		}()
		// 使用Next执行具体的业务逻辑
		_ = c.Next()
		return nil
	}
}
