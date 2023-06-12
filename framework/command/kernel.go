/**
 * @Author: yy
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:14
 */

package command

import "github.com/Joker-desire/go-web/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	// 挂载AppCommand命令
	root.AddCommand(initAppCommand())
}
