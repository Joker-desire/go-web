/**
 * @Author: yy
 * @Description:
 * @File:  simple_response
 * @Version: 1.0.0
 * @Date: 2023/06/07 10:34
 */

package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	// IJson 输出
	IJson(obj any) IResponse

	// IJsonP 输出
	IJsonP(obj any) IResponse

	// IXml 输出
	IXml(obj any) IResponse

	// IHtml 输出
	IHtml(file string, obj any) IResponse

	// IText 输出
	IText(format string, values ...any) IResponse

	// IRedirect 重定向
	IRedirect(path string) IResponse

	// ISetHeader 设置响应头
	ISetHeader(key string, value string) IResponse

	// ISetCookie 设置cookie
	ISetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// ISetStatus 设置状态码
	ISetStatus(code int) IResponse

	// ISetOkStatus 设置200状态码
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj any) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	_, err = ctx.Writer.Write(byt)
	return ctx
}

// IJsonP 输出
func (ctx *Context) IJsonP(obj any) IResponse {
	// 获取请求中的参数作为函数名，获取要返回的数据 JSON 作为函数参数，将函数名 + 函数参数作为返回文本。

	// 获取请求参数callback
	callbackFunc := ctx.Query("callback")
	ctx.ISetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成xss攻击
	callback := template.JSEscapeString(callbackFunc)

	// 输出函数名
	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.Writer.Write([]byte("("))
	if err != nil {
		return ctx
	}

	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	// 输出数据
	_, err = ctx.Writer.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) IXml(obj any) IResponse {
	bytes, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/xml")
	_, err = ctx.Writer.Write(bytes)
	return ctx
}

func (ctx *Context) IHtml(file string, obj any) IResponse {
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
	if err := t.Execute(ctx.Writer, obj); err != nil {
		return ctx
	}
	ctx.ISetHeader("Content-Type", "text/html")
	return ctx
}

func (ctx *Context) IText(format string, values ...any) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type", "application/text")
	_, err := ctx.Writer.Write([]byte(out))
	if err != nil {
		return ctx
	}
	return ctx
}

// IRedirect 重定向
func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

// ISetHeader 设置响应头
func (ctx *Context) ISetHeader(key string, value string) IResponse {
	ctx.Writer.Header().Add(key, value)
	return ctx
}

// ISetCookie 设置cookie
func (ctx *Context) ISetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
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

// ISetStatus 设置状态码
func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

// ISetOkStatus 设置200状态码
func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}
