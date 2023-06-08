/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/07 15:06
 */

package demo

import (
	"fmt"
	"github.com/Joker-desire/go-web/framework"
)

// ServiceProviderDemo ServiceProvider DemoServiceProvider 服务提供方
type ServiceProviderDemo struct {
}

func (s *ServiceProviderDemo) Name() string {
	return Key
}

func (s *ServiceProviderDemo) Register(container framework.Container) framework.NewInstance {
	return NewServiceDemo
}

func (s *ServiceProviderDemo) Params(container framework.Container) []any {
	return []any{container}
}

func (s *ServiceProviderDemo) IsDefer() bool {
	return true
}

// Boot 这里什么逻辑都不执行，只打印一行日志信息
func (s *ServiceProviderDemo) Boot(container framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
