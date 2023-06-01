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
	core.Get("/foo", FooControllerHandler)
	core.Get("/hello", HelloControllerHandler)
	api := core.Group("/api")
	{
		api.Get("/user", UserControllerHandler)
		api.Get("/user/:id/detail", UserControllerHandler2)
		v1 := api.Group("/v1")
		{
			v1.Get("/user", UserControllerHandler)
		}
	}
}
