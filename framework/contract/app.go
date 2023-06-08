/**
 * @Author: yy
 * @Description:
 * @File:  app
 * @Version: 1.0.0
 * @Date: 2023/06/08 12:34
 */

package contract

const AppKey = "go-web:app"

type App interface {
	// Version 定义当前版本号
	Version() string
	// BaseFolder 定义项目基础路径
	BaseFolder() string
	// ConfigFolder 定义配置文件所在路径
	ConfigFolder() string
	// LogFolder 定义日志所在路径
	LogFolder() string
	// ProviderFolder 定义业务自己的服务提供者地址
	ProviderFolder() string
	// MiddlewareFolder 定义业务自定义的中间接地址
	MiddlewareFolder() string
	// CommandFolder 定义业务自定义的命令地址
	CommandFolder() string
	// RuntimeFolder 定义业务的运行中间态信息
	RuntimeFolder() string
	// TestFolder 定义业务的测试文件地址
	TestFolder() string
}
