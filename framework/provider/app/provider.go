/**
 * @Author: yy
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:15
 */

package app

import (
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/contract"
)

type HadeAppProvider struct {
	BaseFolder string
}

func (h *HadeAppProvider) Name() string {
	return contract.AppKey
}

func (h *HadeAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeApp
}

func (h *HadeAppProvider) Params(container framework.Container) []any {
	return []any{container, h.BaseFolder}
}

func (h *HadeAppProvider) IsDefer() bool {
	return false
}

func (h *HadeAppProvider) Boot(container framework.Container) error {
	return nil
}
