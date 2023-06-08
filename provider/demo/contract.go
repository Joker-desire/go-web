/**
 * @Author: yy
 * @Description:
 * @File:  contract
 * @Version: 1.0.0
 * @Date: 2023/06/07 15:06
 */

package demo

// Key Demo 服务的key
const Key = "hade:demo"

// Service Demo服务的接口
type Service interface {
	GetFoo() Foo
}

// Foo Demo服务接口定义的一个数据结构
type Foo struct {
	Name string
}
