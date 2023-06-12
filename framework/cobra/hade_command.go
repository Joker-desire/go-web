/**
 * @Author: yy
 * @Description:
 * @File:  hade_command
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:53
 */

package cobra

import "github.com/Joker-desire/go-web/framework"

// 新增两个方法，设置服务器容器和获取服务容器

// SetContainer 设置服务器容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取服务容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}
