/**
 * @Author: yy
 * @Description:
 * @File:  context
 * @Version: 1.0.0
 * @Date: 2023/06/01 11:46
 */

package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Context 自定义上下文结构
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handler        ControllerHandler

	// 写保护机制
	writerMux *sync.Mutex
	// 是否超时标记位
	hasTimeout bool
}

// NewContext 初始化上下文结构
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      new(sync.Mutex),
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

// request

// QueryInt 获取query int参数
func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		l := len(values)
		if l > 0 {
			intVal, err := strconv.Atoi(values[l-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

// QueryString 获取query string参数
func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		l := len(values)
		if l > 0 {
			return values[l-1]
		}
	}
	return def
}

// QueryArray 获取query array参数
func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// QueryAll 获取query所有参数
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// FormInt 获取form int参数
func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		l := len(values)
		if l > 0 {
			intVal, err := strconv.Atoi(values[l-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

// FormString 获取form string参数
func (ctx *Context) FormString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		l := len(values)
		if l > 0 {
			return values[l-1]
		}
	}
	return def
}

// FormArray 获取form array参数
func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// FormAll 获取form参数
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

// BindJson 将请求体参数与json绑定
func (ctx *Context) BindJson(obj any) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("request is nil")
	}
	return nil
}

// response

// Json 返回json数据
func (ctx *Context) Json(status int, obj any) error {

	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	_, err = ctx.responseWriter.Write(byt)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	return nil
}
