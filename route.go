/**
 * @Author: yy
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2023/06/01 14:05
 */

package main

import "github.com/Joker-desire/go-web/framework"

func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooControllerHandler)
}
