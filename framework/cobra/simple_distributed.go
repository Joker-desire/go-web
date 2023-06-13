/**
 * @Author: yy
 * @Description:
 * @File:  simple_distributed
 * @Version: 1.0.0
 * @Date: 2023/06/12 16:12
 */

package cobra

import (
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// AddDistributedCommand 实现一个分布式定时器
func (c *Command) AddDistributedCommand(serviceName, spec string, cmd *Command, holdTime time.Duration) {
	root := c.Root()

	// 初始化cron
	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}
	//cron命令的注释，这里注意Type为distributed-cron，ServiceName需要填写
	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type:        "distributed-cron",
		Cmd:         cmd,
		Spec:        spec,
		ServiceName: serviceName,
	})

	appService := root.GetContainer().MustMake(contract.AppKey).(contract.App)
	distributedService := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	appID := appService.AppID()

	// 复制要执行的command为cronCmd，并设置为rootCmd
	var cronCmd Command
	ctx := root.Context()
	cronCmd = *cmd
	cronCmd.args = []string{}
	cronCmd.SetParantNull()

	// cron增加匿名函数
	root.Cron.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		// 节点进行选举，返回选举结果
		selectAppID, err := distributedService.Select(serviceName, appID, holdTime)
		if err != nil {
			return
		}

		// 如果自己没有被选中，直接返回
		if selectAppID != appID {
			return
		}
		// 如果自己被选中，执行命令
		err = cronCmd.ExecuteContext(ctx)
		if err != nil {
			log.Println(err)
		}
	})
}
