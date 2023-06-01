/**
 * @Author: yy
 * @Description:
 * @File:  group
 * @Version: 1.0.0
 * @Date: 2023/06/01 16:34
 */

package framework

// IGroup 路由组接口
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

// NewGroup 初始化路由组
func NewGroup(core *Core, prefix string) *Group {
	return &Group{core: core, prefix: prefix}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	g.core.Get(g.prefix+uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	g.core.Get(g.prefix+uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	g.core.Get(g.prefix+uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.core.Get(g.prefix+uri, handler)
}

// Group 实现Group方法
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
