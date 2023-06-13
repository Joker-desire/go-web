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
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/util"
	"github.com/google/uuid"
	"path/filepath"
)

type SimpleApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础文件夹
	appID      string
}

func (h SimpleApp) AppID() string {
	return h.appID
}

func (h SimpleApp) Version() string {
	return "0.0.1"
}

func (h SimpleApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}
	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}
func (h SimpleApp) ConsoleFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "console"))
}

func (h SimpleApp) StorageFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "storage"))
}
func (h SimpleApp) ConfigFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "config"))
}

func (h SimpleApp) LogFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "log"))
}

func (h SimpleApp) ProviderFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "provider"))
}

func (h SimpleApp) MiddlewareFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "middleware"))
}

func (h SimpleApp) CommandFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "command"))
}

func (h SimpleApp) RuntimeFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "runtime"))
}

func (h SimpleApp) TestFolder() string {
	return util.MkdirIfNotExist(filepath.Join(h.BaseFolder(), "test"))
}

func NewSimpleApp(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("params error")
	}
	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
		flag.Parse()
	}
	appID := uuid.New().String()
	return &SimpleApp{container: container, baseFolder: baseFolder, appID: appID}, nil
}
