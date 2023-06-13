/**
 * @Author: yy
 * @Description:
 * @File:  cron
 * @Version: 1.0.0
 * @Date: 2023/06/12 14:59
 */

package command

import (
	"fmt"
	"github.com/Joker-desire/simple/framework/cobra"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/util"
	"github.com/erikdubbelboer/gspt"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var cronDaemon = false

func initCronCommand() *cobra.Command {
	cronStartCommand.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "是否以守护进程方式运行")
	cronCommand.AddCommand(cronStartCommand)
	cronCommand.AddCommand(cronRestartCommand)
	cronCommand.AddCommand(cronStopCommand)
	cronCommand.AddCommand(cronStateCommand)
	cronCommand.AddCommand(cronListCommand)
	return cronCommand
}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var cronListCommand = &cobra.Command{
	Use:   "list",
	Short: "定时任务列表",
	RunE: func(cmd *cobra.Command, args []string) error {
		cronSpecs := cmd.Root().CronSpecs
		var ps [][]string
		for _, cronSpec := range cronSpecs {
			line := []string{
				cronSpec.Type,
				cronSpec.Spec,
				cronSpec.Cmd.Use,
				cronSpec.Cmd.Short,
				cronSpec.ServiceName,
			}
			ps = append(ps, line)
		}
		util.PrettyPrint(ps)
		return nil
	},
}

// cron 进程启动服务
var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 获取容器
		container := cmd.GetContainer()
		// 获取容器中的APP服务
		appService := container.MustMake(contract.AppKey).(contract.App)

		// 设置cron的日志地址和进程ID地址
		pidFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "cron.log")
		currentFolder := appService.BaseFolder()
		// daemon模式
		if cronDaemon {
			// 创建一个Context
			cntxt := &daemon.Context{
				// 设置pid文件
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				// 设置日志文件
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				// 设置工作目录
				WorkDir: currentFolder,
				// 设置所有设置文件的mask，默认为750
				Umask: 027,
				// 子进程的参数，按照这个参数设置，子进程的命令为：simple cron start --daemon=true
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			// 启动子进程，d不为空表示当前是父进程，d为空表示当前是子进程
			d, err := cntxt.Reborn()
			if err != nil {
				fmt.Println("Unable to run: ", err)
			}
			if d != nil {
				// 父进程直接打印启动成功信息，不做任何操作
				fmt.Println("cron serve started pid:", d.Pid)
				fmt.Println("log file:", serverLogFile)
				return nil
			}
			// 子进程执行Cron.Run
			defer cntxt.Release()
			fmt.Println("daemon started")
			gspt.SetProcTitle("simple cron daemon")
			cmd.Root().Cron.Run()
			return nil
		}
		// 非deamon模式
		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(serverPidFile, []byte(content), 0664)
		if err != nil {
			fmt.Println("write pid file error", err)
		}
		gspt.SetProcTitle("simple cron")
		cmd.Root().Cron.Run()
		return nil
	},
}

var cronRestartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重启cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		// 获取PID
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")
		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			fmt.Println("read pid file error", err)
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				fmt.Println("pid file content error", err)
			}
			if util.CheckProcessExist(pid) {
				if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
					return err
				}
				//检查进程是否关闭掉
				for i := 0; i < 10; i++ {
					if util.CheckProcessExist(pid) {
						break
					} else {
						time.Sleep(1 * time.Second)
					}
				}
				fmt.Println("kill process:", pid)
			}
		}
		cronDaemon = true
		return cronStartCommand.RunE(cmd, args)
	},
}

var cronStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "停止cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		// 获取PID
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")
		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			fmt.Println("read pid file error", err)
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				fmt.Println("pid file content error", err)
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := ioutil.WriteFile(serverPidFile, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("stop pid:", pid)
		}
		return nil
	},
}

var cronStateCommand = &cobra.Command{
	Use:   "state",
	Short: "查看cron常驻进程状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		// 获取PID
		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")
		content, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			fmt.Println("read pid file error", err)
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				fmt.Println("pid file content error", err)
			}
			if util.CheckProcessExist(pid) {
				fmt.Println("cron is running, pid:", pid)
				return nil
			}
		}
		return nil
	},
}
