/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/13 12:57
 */

package env

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
)

type SimpleEnvProvider struct {
	Folder string
}

func (s *SimpleEnvProvider) Name() string {
	return contract.EnvKey
}

func (s *SimpleEnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewSimpleEnvService
}

func (s *SimpleEnvProvider) Params(container framework.Container) []any {
	return []any{s.Folder}
}

func (s *SimpleEnvProvider) IsDefer() bool {
	return false
}

func (s *SimpleEnvProvider) Boot(container framework.Container) error {
	app := container.MustMake(contract.AppKey).(contract.App)
	s.Folder = app.BaseFolder()
	return nil
}
