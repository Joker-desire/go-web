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
	"strings"
)

// Core 框架核心结构
type Core struct {
	// 路由表
	//router map[string]ControllerHandler
	router map[string]map[string]ControllerHandler
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	// 定义二级Map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}
	// 将二级Map赋值给一级Map
	router := map[string]map[string]ControllerHandler{
		"GET":    getRouter,
		"POST":   postRouter,
		"PUT":    putRouter,
		"DELETE": deleteRouter,
	}
	return &Core{router: router}
}

// Get 注册GET请求
func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url) // 转换为大写
	c.router["GET"][upperUrl] = handler
}

// Post 注册Post请求
func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url) // 转换为大写
	c.router["POST"][upperUrl] = handler
}

// Put 注册Put请求
func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url) // 转换为大写
	c.router["PUT"][upperUrl] = handler
}

// Delete 注册Delete请求
func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url) // 转换为大写
	c.router["DELETE"][upperUrl] = handler
}

func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// 将URI和Method转换为大写
	uri := request.URL.Path
	method := request.Method
	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	// 根据Method获取对应的Map
	if methodMap, ok := c.router[upperMethod]; ok {
		// 根据URI获取对应的Handler
		if handler, ok := methodMap[upperUri]; ok {
			return handler
		}
	}
	return nil
}

// ServeHTTP 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core serve http")
	ctx := NewContext(response, request)

	// 路由查找
	router := c.FindRouteByRequest(request)
	if router == nil {
		// 如果找不到路由，返回404
		_ = ctx.Json(404, "Not Found")
		return
	}
	log.Println("core router")
	// 调用路由，如果路由执行失败，返回500
	if err := router(ctx); err != nil {
		_ = ctx.Json(500, err.Error())
		return
	}
}
