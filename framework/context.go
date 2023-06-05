/**
 * @Author: yy
 * @Description:
 * @File:  context
 * @Version: 1.0.0
 * @Date: 2023/06/01 11:46
 */

package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// Context 自定义上下文结构
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	//handler        ControllerHandler

	// 写保护机制
	writerMux *sync.Mutex
	// 是否超时标记位
	hasTimeout bool

	// 当前请求的handler链条
	handlers []ControllerHandler
	index    int // 当前请求调用到调用链的哪个节点

	params map[string]string // url路由匹配的参数
}

// NewContext 初始化上下文结构
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      new(sync.Mutex),
		index:          -1, // 初始化值为-1，每次调用都会自增 ，这样首次调用的时候index=0
	}
}

// base

// WriterMux 对外暴露锁
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

// GetRequest 对外暴露request
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

// GetResponse 对外暴露response
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

// SetHasTimeout 设置超时标记位
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

// HasTimeout 获取超时标记位
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// SetHandlers 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

// Next 核心函数，调用context的下一个函数
// 为了控制实现链条的逐步调用，需要在每个中间件中调用Next函数。
// 这个Next方法每调用一次，讲讲这个控制器链路的调用控制器
// 通过Next函数的调用，可以实现链条的逐步调用，直到调用到最后一个中间件。
// 通过index下标表示当前调用Next要执行的控制器序列
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// context

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.request.Context().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.request.Context().Done()

}

func (ctx *Context) Err() error {
	return ctx.request.Context().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.request.Context().Value(key)
}
