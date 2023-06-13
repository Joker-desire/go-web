/**
 * @Author: yy
 * @Description:
 * @File:  simple_engine
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:57
 */

package gin

import "github.com/Joker-desire/simple/framework"

//engine实现container的绑定封装

// Bind 关键字凭证绑定服务提供者
func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind 关键字凭证是否已经绑定服务提供者
func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

func (engine *Engine) SerContainer(container framework.Container) {
	engine.container = container
}
