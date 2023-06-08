/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:30
 */

package app

import (
	"errors"
	"flag"
	"github.com/Joker-desire/go-web/framework"
	"github.com/Joker-desire/go-web/framework/util"
	"path/filepath"
)

type HadeApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础文件夹
}

func (h HadeApp) Version() string {
	return "0.0.1"
}

func (h HadeApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}
func (h HadeApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(), "console")
}

func (h HadeApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(), "storage")
}
func (h HadeApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

func (h HadeApp) LogFolder() string {
	return filepath.Join(h.BaseFolder(), "log")
}

func (h HadeApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func (h HadeApp) MiddlewareFolder() string {
	return filepath.Join(h.BaseFolder(), "middleware")
}

func (h HadeApp) CommandFolder() string {
	return filepath.Join(h.BaseFolder(), "command")
}

func (h HadeApp) RuntimeFolder() string {
	return filepath.Join(h.BaseFolder(), "runtime")
}

func (h HadeApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "test")
}

func NewHadeApp(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}
	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{container: container, baseFolder: baseFolder}, nil
}
