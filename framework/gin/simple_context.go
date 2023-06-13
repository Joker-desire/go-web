/**
 * @Author: yy
 * @Description:
 * @File:  simple_context
 * @Version: 1.0.0
 * @Date: 2023/06/07 10:54
 */

package gin

import (
	"context"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
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
