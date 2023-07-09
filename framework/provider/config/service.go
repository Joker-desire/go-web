/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/30 17:16
 */

package config

import (
	"bytes"
	"errors"
	"github.com/Joker-desire/simple/framework"
	"github.com/go-yaml/yaml"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type SimpleConfig struct {
	c        framework.Container // 容器
	folder   string              // 文件夹
	keyDelim string              // 路径分隔符，默认为.
	envMaps  map[string]string   // 环境变量
	confMaps map[string]any      // 配置文件结构，key为文件名
	confRaws map[string][]byte   // 配置文件原始内容
	lock     sync.RWMutex        // 配置文件读写锁
}

func NewSimpleConfig(params ...any) (any, error) {
	container := params[0].(framework.Container)
	envFolder := params[1].(string)
	envMaps := params[2].(map[string]string)

	// 检查文件夹是否存在
	if _, err := os.Stat(envFolder); os.IsNotExist(err) {
		return nil, errors.New("folder" + envFolder + "not exist: " + err.Error())
	}

	// 实例化
	simpleConf := &SimpleConfig{
		c:        container,
		folder:   envFolder,
		keyDelim: ".",
		envMaps:  envMaps,
		confMaps: make(map[string]any),
		confRaws: make(map[string][]byte),
		lock:     sync.RWMutex{},
	}
	// 读取文件夹下的所有文件
	files, err := ioutil.ReadDir(envFolder)
	if err != nil {
		return nil, err
	}
	// 读取文件内容
	for _, file := range files {
		fileName := file.Name()
		err := simpleConf.loadConfigFile(envFolder, fileName)
		if err != nil {
			log.Println("load config file error: ", err.Error())
			continue
		}
	}
	return simpleConf, nil
}

func (c *SimpleConfig) loadConfigFile(folder, file string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	// 判断文件是否以yaml或者yml作为后缀
	split := strings.Split(file, ".")
	if len(split) == 2 && (split[1] == "yaml" || split[1] == "yml") {
		name := split[0]
		// 读取文件内容
		bf, err := ioutil.ReadFile(filepath.Join(folder, file))
		if err != nil {
			return err
		}
		// 直接针对文本做环境变量的替换
		bf = replace(bf, c.envMaps)
		// 解析对应的文件
		confMap := map[string]any{}
		if err := yaml.Unmarshal(bf, &confMap); err != nil {
			return err
		}
		c.confMaps[name] = confMap
		c.confRaws[name] = bf
	}
	return nil
}

// replace 表示使用环境变量maps替换content中的环境变量（env(xxx)）
func replace(content []byte, maps map[string]string) []byte {
	if maps == nil {
		return content
	}
	// 替换配置文件中的环境变量
	for k, v := range maps {
		reKey := "env(" + k + ")"
		content = bytes.ReplaceAll(content, []byte(reKey), []byte(v))
	}
	return content
}

// 查找某个路径的配置项
func searchMap(source map[string]any, path []string) any {

	if len(path) == 0 {
		return source
	}

	// 递归查找
	next, ok := source[path[0]]
	if ok {
		// 判断这个路径是否为1
		if len(path) == 1 {
			return next
		}
		// 判断下一个路径的类型
		switch next.(type) {
		case map[any]any:
			// 如果是any的Map，使用cast进行下value转换
			return searchMap(cast.ToStringMap(next), path[1:])
		case map[string]any:
			// 如果是string的Map，直接进行下一次递归
			return searchMap(next.(map[string]any), path[1:])
		default:
			// 否则返回nil
			return nil
		}
	}
	return nil
}

func (c *SimpleConfig) find(key string) any {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return searchMap(c.confMaps, strings.Split(key, c.keyDelim))
}

func (c *SimpleConfig) IsExist(s string) bool {
	return c.find(s) != nil
}

func (c *SimpleConfig) Get(s string) any {
	return c.find(s)
}

func (c *SimpleConfig) GetInt(s string) int {
	return cast.ToInt(c.find(s))
}

func (c *SimpleConfig) GetFloat64(s string) float64 {
	return cast.ToFloat64(c.find(s))
}

func (c *SimpleConfig) GetString(s string) string {
	return cast.ToString(c.find(s))
}

func (c *SimpleConfig) GetBool(s string) bool {
	return cast.ToBool(c.find(s))
}

func (c *SimpleConfig) GetTime(s string) time.Time {
	return cast.ToTime(c.find(s))
}

func (c *SimpleConfig) GetIntSlice(s string) []int {
	return cast.ToIntSlice(c.find(s))
}

func (c *SimpleConfig) GetStringSlice(s string) []string {
	return cast.ToStringSlice(c.find(s))
}

func (c *SimpleConfig) GetStringMap(s string) map[string]any {
	return cast.ToStringMap(c.find(s))
}

func (c *SimpleConfig) GetStringMapString(s string) map[string]string {
	return cast.ToStringMapString(c.find(s))
}

func (c *SimpleConfig) GetStringMapStringSlice(s string) map[string][]string {
	return cast.ToStringMapStringSlice(c.find(s))
}

func (c *SimpleConfig) Load(key string, val any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  val,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(c.find(key))
}
