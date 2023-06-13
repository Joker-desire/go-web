/**
 * @Author: yy
 * @Description:
 * @File:  exec
 * @Version: 1.0.0
 * @Date: 2023/06/08 14:37
 */

package util

import (
	"os"
	"syscall"
)

// GetExecDirectory 获取当前执行文件的目录
func GetExecDirectory() string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}
	return file + "/"
}

// CheckProcessExist 检查进程是否存在
func CheckProcessExist(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}
