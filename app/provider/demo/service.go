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

// Service 具体的接口实例
type Service struct {
	// 实现接口
	IService
	// 参数
	c framework.Container
}

// NewService 实例化接口
func NewService(params ...any) (any, error) {
	// 这里需要将参数展开
	c := params[0].(framework.Container)
	fmt.Println("new demo service")
	// 返回实例
	return &Service{c: c}, nil
}

// GetAllStudent 实现接口的方法
func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "yy",
		},
		{
			ID:   2,
			Name: "Joker",
		},
	}
}
