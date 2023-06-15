/**
 * @Author: yy
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2023/06/13 12:32
 */

package contract

const (
	EnvProduction  = "production"  // 生产环境
	EnvDevelopment = "development" // 开发环境
	EnvKey         = "simple:env"  // 环境变量服务字符串凭证
)

type Env interface {
	// AppEnv 返回当前的环境变量
	AppEnv() string
	// IsExist 判断某个环境变量是否存在
	IsExist(string) bool
	// Get 获取某个环境变量的值
	Get(string) string
	// All 获取所有的环境变量
	All() map[string]string
}
