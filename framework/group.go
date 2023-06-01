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

	// Group 实现嵌套group
	Group(string) IGroup
}

type Group struct {
	core   *Core  //指向core结构
	parent *Group // 指向上一个Group，如果有的话
	prefix string //路由前缀
}

// NewGroup 初始化路由组
func NewGroup(core *Core, prefix string) *Group {
	return &Group{core: core, parent: nil, prefix: prefix}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	g.core.Get(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	g.core.Get(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	g.core.Get(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.core.Get(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

// Group 实现Group方法
func (g *Group) Group(prefix string) IGroup {
	cGroup := NewGroup(g.core, prefix)
	cGroup.parent = g
	return cGroup
}
