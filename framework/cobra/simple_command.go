/**
 * @Author: yy
 * @Description:
 * @File:  simple_command
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:53
 */

package cobra

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/robfig/cron/v3"
	"log"
)

// 新增两个方法，设置服务器容器和获取服务容器

// SetContainer 设置服务器容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取服务容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}

// CronSpec 保存Cron命令的信息，用于展示
type CronSpec struct {
	Type        string
	Cmd         *Command
	Spec        string
	ServiceName string
}

func (c *Command) SetParantNull() {
	c.parent = nil
}

// AddCronCommand 用来创建一个Cron命令
func (c *Command) AddCronCommand(spec string, cmd *Command) {
	// cron结构是挂载在根Command上的
	root := c.Root()
	if root.Cron == nil {
		// 初始化cron
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	// 增加说明信息
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type: "normal-cron",
		Cmd:  cmd,
		Spec: spec,
	})
	// 只做一个rootCommand
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParantNull()
	cronCmd.SetContainer(root.GetContainer())

	// 增加调用函数
	root.Cron.AddFunc(spec, func() {
		// 如果后缀的command出现panic,这里进行捕获
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		err := cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Println(err)
		}
	})
}
