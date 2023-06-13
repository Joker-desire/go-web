/**
 * @Author: yy
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2023/06/02 14:37
 */

package middleware

import (
	"github.com/Joker-desire/simple/framework/gin"
	"log"
)

func Test1() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		log.Println("middleware pre test1")
		c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test1")
	}
}
func Test2() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		log.Println("middleware pre test2")
		c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test2")
	}
}
func Test3() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		log.Println("middleware pre test3")
		c.Next() // 通过Next往下调用，会自增context.index
		log.Println("middleware post test3")
	}
}
