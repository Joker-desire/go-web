/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:44
 */

package kernel

import (
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/contract"
	"github.com/Joker-desire/go-web/framework/gin"
)

// HadeKernelProvider 提供web引擎
type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *HadeKernelProvider) Name() string {
	return contract.KernelKey
}

func (p *HadeKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeKernelService
}

func (p *HadeKernelProvider) Params(container framework.Container) []any {
	return []any{p.HttpEngine}
}

func (p *HadeKernelProvider) IsDefer() bool {
	return false
}

func (p *HadeKernelProvider) Boot(container framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SerContainer(container)
	return nil
}
