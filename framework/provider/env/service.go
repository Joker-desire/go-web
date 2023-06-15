/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/13 12:35
 */

package env

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/Joker-desire/simple/framework/contract"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type SimpleEnvService struct {
	folder string            // 表示.env所在的目录
	maps   map[string]string // 保存所有的环境变量
}

func (s *SimpleEnvService) AppEnv() string {
	return s.Get("APP_ENV")
}

func (s *SimpleEnvService) IsExist(key string) bool {
	_, ok := s.maps[key]
	return ok
}

func (s *SimpleEnvService) Get(key string) string {
	if val, ok := s.maps[key]; ok {
		return val
	}
	return ""
}

func (s *SimpleEnvService) All() map[string]string {
	return s.maps
}

// NewSimpleEnvService 创建一个新的环境变量服务
func NewSimpleEnvService(params ...any) (any, error) {
	if len(params) != 1 {
		return nil, errors.New("NewSimpleEnvService params error")
	}

	// 读取folder文件
	folder := params[0].(string)

	// 实例化
	simpleEnv := &SimpleEnvService{
		folder: folder,
		maps:   map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	fmt.Println("simpleEnv:", simpleEnv.maps)

	// 解析folder/.env文件
	file := path.Join(folder, ".env")
	log.Println("env file:", file)

	// 打开文件.env
	if fi, err := os.Open(file); err == nil {
		defer fi.Close()

		// 读取文件
		reader := bufio.NewReader(fi)
		for {
			// 按照行进行读取
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			// 按照等号解析
			s := bytes.SplitN(line, []byte{'='}, 2)
			fmt.Println(s)
			// 如果不符合规范，则过滤
			if len(s) < 2 {
				continue
			}
			// 保存到maps中
			key := string(s[0])
			val := string(s[1])
			simpleEnv.maps[key] = val
		}
	}
	fmt.Println("simpleEnv:", simpleEnv.maps)

	// 获取当前程序的环境变量，并且覆盖.env文件下的变量
	for _, e := range os.Environ() {
		n := strings.SplitN(e, "=", 2)
		if len(n) < 2 {
			continue
		}
		simpleEnv.maps[n[0]] = n[1]
	}
	return simpleEnv, nil
}
