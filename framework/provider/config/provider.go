/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/30 17:39
 */

package config

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
	"path/filepath"
)

type SimpleConfigProvider struct {
}

func (s *SimpleConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewSimpleConfig
}

func (s *SimpleConfigProvider) Name() string {
	return contract.ConfigKey
}

func (s *SimpleConfigProvider) Params(container framework.Container) []any {
	appService := container.MustMake(contract.AppKey).(contract.App)
	envService := container.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	// 配置文件夹地址
	configFolder := appService.ConfigFolder()
	envFolder := filepath.Join(configFolder, env)
	return []any{container, envFolder, envService.All()}
}

func (s *SimpleConfigProvider) IsDefer() bool {
	return false
}

func (s *SimpleConfigProvider) Boot(container framework.Container) error {
	return nil
}
