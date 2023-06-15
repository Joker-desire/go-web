/**
 * @Author: yy
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:14
 */

package command

import "github.com/Joker-desire/simple/framework/cobra"

func AddKernelCommands(root *cobra.Command) {

	// 挂载EnvCommand命令
	root.AddCommand(initEnvCommand())

	// 挂载CronCommand命令
	root.AddCommand(initCronCommand())

	// 挂载AppCommand命令
	root.AddCommand(initAppCommand())
}
