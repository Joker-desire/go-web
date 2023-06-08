/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/07 15:06
 */

package demo

import (
	"fmt"
	"github.com/Joker-desire/go-web/framework"
)

// ServiceDemo 具体的接口实例
type ServiceDemo struct {
	// 实现接口
	Service
	// 参数
	c framework.Container
}

// NewServiceDemo 实例化接口
func NewServiceDemo(params ...any) (any, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)
	fmt.Println("new demo service")
	// 返回实例
	return &ServiceDemo{c: c}, nil
}

// GetFoo 实现接口的方法
func (s *ServiceDemo) GetFoo() Foo {
	return Foo{
		Name: "i am demo service foo",
	}
}
