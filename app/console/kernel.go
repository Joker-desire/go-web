/**
 * @Author: yy
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2023/06/08 12:46
 */

package console

import (
	"github.com/Joker-desire/simple/app/console/command/demo"
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/cobra"
	"github.com/Joker-desire/simple/framework/command"
	"time"
)

// RunCommand 初始化根Command并运行
func RunCommand(container framework.Container) error {
	//根Command
	var rootCmd = &cobra.Command{
		// 定义根命令的关键字
		Use: "simple",
		// 定义根命令的简介
		Short: "simple 命令",
		// 定义根命令的详细介绍
		Long: `simple is a web framework`,
		// 定义根命令的运行函数
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		// 不需要出现cobra默认的completion子命令
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	// 为根Command设置服务容器
	rootCmd.SetContainer(container)
	// 绑定框架的命令
	command.AddKernelCommands(rootCmd)
	// 绑定业务的命令
	AddAppCommand(rootCmd)

	// 执行RootCommand
	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	// 挂载AppCommand命令
	rootCmd.AddCommand(demo.InitFoo())

	//rootCmd.AddCronCommand("* * * * * *", demo.FooCommand)
	rootCmd.AddDistributedCommand("foo_func_for_test", "*/5 * * * * *", demo.FooCommand, 2*time.Second)

}
