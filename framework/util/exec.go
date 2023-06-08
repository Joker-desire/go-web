/**
 * @Author: yy
 * @Description:
 * @File:  exec
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:37
 */

package util

import "os"

// GetExecDirectory 获取当前执行文件的目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}
	return file + "/"
}
