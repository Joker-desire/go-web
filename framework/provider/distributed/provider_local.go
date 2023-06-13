/**
 * @Author: yy
 * @Description:
 * @File:  provider_local
 * @Version: 1.0.0
 * @Date: 2023/06/12 15:58
 */

package distributed

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
)

type LocalDistributedProvider struct {
}

func (l LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}

func (l LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

func (l LocalDistributedProvider) Params(container framework.Container) []any {
	return []any{container}
}

func (l LocalDistributedProvider) IsDefer() bool {
	return false
}

func (l LocalDistributedProvider) Boot(container framework.Container) error {
	return nil
}
