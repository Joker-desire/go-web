/**
 * @Author: yy
 * @Description:
 * @File:  hade_request
 * @Version: 1.0.0
 * @Date: 2023/06/06 18:06
 */

package gin

import (
	"github.com/spf13/cast"
)

// IRequest 请求接口
type IRequest interface {

	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar

	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id

	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamString(key string, def string) (string, bool)

	// form表单中带的参数

	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultForm(key string) any // json body中的参数

}

var _ IRequest = &Context{}

// QueryAll 获取DefaultQuery所有参数
func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return map[string][]string(ctx.queryCache)
}

// DefaultQueryInt 获取DefaultQuery int参数
func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

// DefaultQueryInt64 获取DefaultQuery int64参数
func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

// DefaultQueryFloat32 获取DefaultQuery float32参数
func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

// DefaultQueryFloat64 获取DefaultQuery float64参数
func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}

// DefaultQueryString 获取DefaultQuery string参数
func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0], true
		}
	}
	return def, false
}

// DefaultQueryBool 获取DefaultQuery bool参数
func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

// DefaultQueryStringSlice 获取DefaultQuery string slice参数
func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return cast.ToStringSlice(values), true
	}
	return def, false
}

// 获取路由参数
func (ctx *Context) HadeParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

// 路由匹配中带的参数

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if value := ctx.HadeParam(key); value != nil {
		return cast.ToInt(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if value := ctx.HadeParam(key); value != nil {
		return cast.ToInt64(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if value := ctx.HadeParam(key); value != nil {
		return cast.ToFloat32(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if value := ctx.HadeParam(key); value != nil {
		return cast.ToFloat64(value), true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if value := ctx.HadeParam(key); value != nil {
		return cast.ToString(value), true
	}
	return def, false
}

// FormAll 获取DefaultForm参数
func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return map[string][]string(ctx.formCache)
}

// DefaultForm 获取DefaultForm参数
func (ctx *Context) DefaultForm(key string) any {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0]
		}
	}
	return nil
}

// DefaultFormInt 获取DefaultForm int参数
func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt(values[0]), true
		}
	}
	return def, false
}

// DefaultFormInt64 获取DefaultForm int64参数
func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToInt64(values[0]), true
		}
	}
	return def, false
}

// DefaultFormFloat32 获取DefaultForm float32参数
func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat32(values[0]), true
		}
	}
	return def, false
}

// DefaultFormFloat64 获取DefaultForm float64参数
func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToFloat64(values[0]), true
		}
	}
	return def, false
}
func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return cast.ToBool(values[0]), true
		}
	}
	return def, false
}

// DefaultFormString 获取DefaultForm string参数
func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		if len(values) > 0 {
			return values[0], true
		}
	}
	return def, false
}

// DefaultFormStringSlice 获取DefaultForm string slice参数
func (ctx *Context) DefaultFormStringSlice(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}
