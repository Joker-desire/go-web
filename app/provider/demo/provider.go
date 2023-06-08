/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/07 15:06
 */

package demo

import (
	"github.com/Joker-desire/go-web/framework"
)

// ServiceProviderDemo ServiceProvider DemoServiceProvider 服务提供方
type ServiceProviderDemo struct {
	framework.ServiceProvider

	c framework.Container
}

func (s *ServiceProviderDemo) Name() string {
	return Key
}

func (s *ServiceProviderDemo) Register(container framework.Container) framework.NewInstance {
	return NewService
}

func (s *ServiceProviderDemo) Params(container framework.Container) []any {
	return []any{s.c}
}

func (s *ServiceProviderDemo) IsDefer() bool {
	return true
}

func (s *ServiceProviderDemo) Boot(container framework.Container) error {
	s.c = container
	return nil
}
