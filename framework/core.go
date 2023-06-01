/**
 * @Author: yy
 * @Description:
 * @File:  core
 * @Version: 1.0.0
 * @Date: 2023/06/01 10:26
 */

package framework

import "net/http"

// Core 框架核心结构
type Core struct {
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	return &Core{}
}

// ServeHTTP 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
