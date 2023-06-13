/**
 * @Author: yy
 * @Description:
 * @File:  foo
 * @Version: 1.0.0
 * @Date: 2023/06/12 12:34
 */

package demo

import (
	"github.com/Joker-desire/simple/framework/cobra"
	"log"
)

// InitFoo 初始化foo命令
func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo命令",
	Long:    `foo命令`,
	Aliases: []string{"f", "fo"},
	Example: "foo命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		log.Println("foo")
		return nil
	},
}

var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1命令",
	Long:    `foo1命令`,
	Aliases: []string{"f1", "fo1"},
	Example: "foo1命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		log.Println("foo1")
		return nil
	},
}
