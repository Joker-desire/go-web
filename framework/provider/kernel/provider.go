/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:44
 */

package kernel

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/gin"
)

// SimpleKernelProvider 提供web引擎
type SimpleKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *SimpleKernelProvider) Name() string {
	return contract.KernelKey
}

func (p *SimpleKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewSimpleKernelService
}

func (p *SimpleKernelProvider) Params(container framework.Container) []any {
	return []any{p.HttpEngine}
}

func (p *SimpleKernelProvider) IsDefer() bool {
	return false
}

func (p *SimpleKernelProvider) Boot(container framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SerContainer(container)
	return nil
}
