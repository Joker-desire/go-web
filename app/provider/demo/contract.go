/**
 * @Author: yy
 * @Description:
 * @File:  contract
 * @Version: 1.0.0
 * @Date: 2023/06/07 15:06
 */

package demo

// Key Demo 服务的key
const Key = "demo"

// IService Demo服务的接口
type IService interface {
	GetAllStudent() []Student
}

// Student Demo服务接口定义的一个数据结构
type Student struct {
	ID   int
	Name string
}
