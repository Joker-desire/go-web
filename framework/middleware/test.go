/**
 * @Author: yy
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2023/06/02 14:37
 */

package middleware

import (
	"github.com/Joker-desire/go-web/framework"
	"log"
)

func Test1() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		log.Println("middleware pre test1")
		_ = c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test1")
		return nil
	}
}
func Test2() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		log.Println("middleware pre test2")
		_ = c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test2")
		return nil
	}
}
func Test3() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		log.Println("middleware pre test3")
		_ = c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test3")
		return nil
	}
}
