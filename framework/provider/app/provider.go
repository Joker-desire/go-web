/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:15
 */

package app

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
)

type SimpleAppProvider struct {
	BaseFolder string
}

func (h *SimpleAppProvider) Name() string {
	return contract.AppKey
}

func (h *SimpleAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewSimpleApp
}

func (h *SimpleAppProvider) Params(container framework.Container) []any {
	return []any{container, h.BaseFolder}
}

func (h *SimpleAppProvider) IsDefer() bool {
	return false
}

func (h *SimpleAppProvider) Boot(container framework.Container) error {
	return nil
}
