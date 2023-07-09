/**
 * @Author: yy
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2023/06/30 16:23
 */

package contract

import "time"

const ConfigKey = "simple:config"

// Config 定义了配置文件服务，读取配置文件，支持点分割的路径读取
// 例如：.Get("app.name")，表示读取app.yaml文件中的name属性
// 建议配置文件使用yaml格式
type Config interface {
	// IsExist 检查一个属性是否存在
	IsExist(string) bool
	// Get 获取一个属性的值
	Get(string) any
	// GetInt 获取一个int类型的属性值
	GetInt(string) int
	// GetFloat64 获取一个float64类型的属性值
	GetFloat64(string) float64
	// GetString 获取一个string类型的属性值
	GetString(string) string
	// GetBool 获取一个bool类型的属性值
	GetBool(string) bool
	// GetTime 获取一个time.Time类型的属性值
	GetTime(string) time.Time
	// GetIntSlice 获取一个[]int类型的属性值
	GetIntSlice(string) []int
	// GetStringSlice 获取一个[]string类型的属性值
	GetStringSlice(string) []string
	// GetStringMap 获取一个map[string]any类型的属性值
	GetStringMap(string) map[string]any
	// GetStringMapString 获取一个map[string]string类型的属性值
	GetStringMapString(string) map[string]string
	// GetStringMapStringSlice 获取一个map[string][]string类型的属性值
	GetStringMapStringSlice(string) map[string][]string
	// Load 加载配置到某个对象
	Load(key string, val any) error
}
