/**
 * @Author: yy
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2023/06/01 10:26
 */

package framework

import (
	"log"
	"net/http"
)

// Core 框架核心结构
type Core struct {
	// 路由表
	router map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// ServeHTTP 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core serve http")
	ctx := NewContext(response, request)

	// 路由匹配(这里写死foo)
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core router")
	if err := router(ctx); err != nil {
		panic(err)
	}
}
