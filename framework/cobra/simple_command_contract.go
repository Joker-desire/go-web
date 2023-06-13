/**
 * @Author: yy
 * @Description:
 * @File:  simple_command_contract
 * @Version: 1.0.0
 * @Date: 2023/06/12 16:08
 */

package cobra

import "github.com/Joker-desire/simple/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Command) MustMakeApp() contract.App {
	return c.GetContainer().MustMake(contract.AppKey).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Command) MustMakeKernel() contract.Kernel {
	return c.GetContainer().MustMake(contract.KernelKey).(contract.Kernel)
}
