/**
 * @Author: yy
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 2023/06/02 17:52
 */

package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"mime/multipart"
)

const defaultMultipartMemory = 32 << 20 // 32 MB

// IRequest 请求接口
type IRequest interface {

	// 请求地址URL中带的参数(/user?id=1)

	Query(key string) any
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	QueryBool(key string, def bool) (bool, bool)

	// 路由匹配中带的参数（/user/:id）

	Param(key string) any
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamString(key string, def string) (string, bool)

	// 表单中的参数

	Form(key string) any
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) []string
	FormFile(key string) (*multipart.FileHeader, error)

	// json body中的参数

	BindJson(obj any) error

	// xml body中的参数

	BindXml(obj any) error

	// 其他格式

	GetRawData() ([]byte, error)

	// 基础信息

	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// 请求头

	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie信息

	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// QueryAll 获取query所有参数
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// Query 获取query参数
func (ctx *Context) Query(key string) any {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return nil
}

// QueryInt 获取query int参数
func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

// QueryInt64 获取query int64参数
func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

// QueryFloat32 获取query float32参数
func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

// QueryFloat64 获取query float64参数
func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

// QueryString 获取query string参数
func (ctx *Context) QueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0], true
		}
	}
	return def, false
}

// QueryBool 获取query bool参数
func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

// QueryStringSlice 获取query string slice参数
func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return cast.ToStringSlice(values), true
	}
	return def, false
}

func (ctx *Context) Param(key string) any {
	if ctx.params != nil {
		if val, ok := ctx.params[key]; ok {
			return val
		}
	}
	return nil
}

func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if value := ctx.Param(key); value != nil {
		return cast.ToInt(value), true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if value := ctx.Param(key); value != nil {
		return cast.ToInt64(value), true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if value := ctx.Param(key); value != nil {
		return cast.ToFloat32(value), true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if value := ctx.Param(key); value != nil {
		return cast.ToFloat64(value), true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if value := ctx.Param(key); value != nil {
		return cast.ToString(value), true
	}
	return def, false
}

// FormAll 获取form参数
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

// Form 获取form参数
func (ctx *Context) Form(key string) any {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return nil
}

// FormInt 获取form int参数
func (ctx *Context) FormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

// FormInt64 获取form int64参数
func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

// FormFloat32 获取form float32参数
func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

// FormFloat64 获取form float64参数
func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

// FormString 获取form string参数
func (ctx *Context) FormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0], true
		}
	}
	return def, false
}

// FormStringSlice 获取form string slice参数
func (ctx *Context) FormStringSlice(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

// FormFile 获取form file参数
func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {

	if ctx.request.MultipartForm == nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	file, header, err := ctx.request.FormFile(key)
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return header, err

}

// BindJson 将请求体参数与json绑定
func (ctx *Context) BindJson(obj any) error {
	if ctx.request != nil {
		// 读取请求体，ctx.request.Body的读取是一次性的
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}

		// 重新填充request.Body，为后续的逻辑二次读取做准备
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 解析到obj结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("request is nil")
	}
	return nil
}

// BindXml 将请求体参数与xml绑定
func (ctx *Context) BindXml(obj any) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("request is nil")
	}
	return nil
}

// GetRawData 获取请求体原始数据
func (ctx *Context) GetRawData() ([]byte, error) {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return nil, err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, errors.New("request is nil")
}

func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string {
	ipAddr := ctx.request.Header.Get("X-Real-Ip")
	if ipAddr == "" {
		ipAddr = ctx.request.Header.Get("X-Forwarded-For")
	}
	if ipAddr == "" {
		ipAddr = ctx.request.RemoteAddr
	}
	return ipAddr
}

func (ctx *Context) Headers() map[string][]string {
	return ctx.request.Header
}

func (ctx *Context) Header(key string) (string, bool) {
	values := ctx.request.Header.Values(key)
	if values == nil || len(values) <= 0 {
		return "", false
	}
	return values[0], true
}

func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if value, ok := cookies[key]; ok {
		return value, true
	}
	return "", false
}
