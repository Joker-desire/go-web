/**
 * @Author: yy
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2023/06/13 12:50
 */

package command

import (
	"fmt"
	"github.com/Joker-desire/simple/framework/cobra"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/util"
)

func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的App环境",
	Run: func(cmd *cobra.Command, args []string) {
		//获取env环境变量
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		// 打印环境
		fmt.Println("env:", envService.AppEnv())
	},
}

var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取所有的环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		//获取env环境变量
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		var outs [][]string
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
