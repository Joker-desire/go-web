/**
 * @Author: yy
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2023/06/02 17:52
 */

package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	// Json 输出
	Json(obj any) IResponse

	// JsonP 输出
	JsonP(obj any) IResponse

	// Xml 输出
	Xml(obj any) IResponse

	// Html 输出
	Html(file string, obj any) IResponse

	// Text 输出
	Text(format string, values ...any) IResponse

	// Redirect 重定向
	Redirect(path string) IResponse

	// SetHeader 设置响应头
	SetHeader(key string, value string) IResponse

	// SetCookie 设置cookie
	SetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// SetStatus 设置状态码
	SetStatus(code int) IResponse

	// SetOkStatus 设置200状态码
	SetOkStatus() IResponse
}

func (ctx *Context) Json(obj any) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	_, err = ctx.responseWriter.Write(byt)
	return ctx
}

// JsonP 输出
func (ctx *Context) JsonP(obj any) IResponse {
	// 获取请求中的参数作为函数名，获取要返回的数据 JSON 作为函数参数，将函数名 + 函数参数作为返回文本。

	// 获取请求参数callback
	callbackFunc, _ := ctx.QueryString("callback", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	// 输出数据
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) Xml(obj any) IResponse {
	bytes, err := xml.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/xml")
	_, err = ctx.responseWriter.Write(bytes)
	return ctx
}

func (ctx *Context) Html(file string, obj any) IResponse {
	/**
	1. 先根据模版创造出 template 结构；
	2. 再使用 template.Execute 将传入数据和模版结合。
	*/
	// 读取模板文件
	t, err := template.ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 执行Execute方法，将数据写入模板
	if err := t.Execute(ctx.responseWriter, obj); err != nil {
		return ctx
	}
	ctx.SetHeader("Content-Type", "text/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...any) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.SetHeader("Content-Type", "application/text")
	_, err := ctx.responseWriter.Write([]byte(out))
	if err != nil {
		return ctx
	}
	return ctx
}

// Redirect 重定向
func (ctx *Context) Redirect(path string) IResponse {
	http.Redirect(ctx.responseWriter, ctx.request, path, http.StatusMovedPermanently)
	return ctx
}

// SetHeader 设置响应头
func (ctx *Context) SetHeader(key string, value string) IResponse {
	ctx.responseWriter.Header().Add(key, value)
	return ctx
}

// SetCookie 设置cookie
func (ctx *Context) SetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.responseWriter, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

// SetStatus 设置状态码
func (ctx *Context) SetStatus(code int) IResponse {
	ctx.responseWriter.WriteHeader(code)
	return ctx
}

// SetOkStatus 设置200状态码
func (ctx *Context) SetOkStatus() IResponse {
	ctx.responseWriter.WriteHeader(http.StatusOK)
	return ctx
}
