/**
 * @Author: yy
 * @Description:
 * @File:  hade_context
 * @Version: 1.0.0
 * @Date: 2023/06/07 10:54
 */

package gin

import (
	"context"
	"github.com/Joker-desire/go-web/framework"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}

//engine实现container的绑定封装

// Bind 关键字凭证绑定服务提供者
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 实现container的make封装

// Make 实现make的封装
func (ctx *Context) Make(key string) (any, error) {
	return ctx.container.Make(key)
}

// MustMake 实现MustMake的封装
func (ctx *Context) MustMake(key string) any {
	return ctx.container.MustMake(key)
}

// MakeNew 实现MakeNew的封装
func (ctx *Context) MakeNew(key string, params []any) (any, error) {
	return ctx.container.MakeNew(key, params)
}
