/**
 * @Author: yy
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/06/12 16:48
 */

package util

import "os"

// MkdirIfNotExist 如果文件夹不存在，创建文件夹
func MkdirIfNotExist(path string) string {
	_, err := os.Stat(path)
	if err != nil {
		os.MkdirAll(path, os.ModePerm)
	}
	return path
}
