/**
 * @Author: yy
 * @Description: 服务容器
 * @File:  container
 * @Version: 1.0.0
 * @Date: 2023/06/07 12:35
 */

package framework

import (
	"errors"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，不返回error
	Bind(provider ServiceProvider) error
	// IsBind 关键字凭证是否已经绑定服务提供者
	IsBind(key string) bool

	// Make 提供关键字凭证获取一个服务
	Make(key string) (any, error)
	// MustMake 提供关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，会panic
	// 所以在使用这个接口的时候清保证服务容器已经为这个关键字凭证绑定了服务提供者
	MustMake(key string) any
	// MakeNew 提供关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的params参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	MakeNew(key string, params []any) (any, error)
}

type SimpleContainer struct {
	Container                            // 强制要求SimpleContainer实现Container接口
	providers map[string]ServiceProvider // 存储注册的服务提供者，key是服务提供者的凭证
	instances map[string]any             // 存储具体的实例，key是服务提供者的凭证
	lock      sync.RWMutex               // 用于锁住对容器的变更操作
}

// NewSimpleContainer 创建一个服务容器
func NewSimpleContainer() *SimpleContainer {
	return &SimpleContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]any{},
		lock:      sync.RWMutex{},
	}
}

// 查找一个服务提供者
func (h *SimpleContainer) findServiceProvider(key string) ServiceProvider {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if provider, ok := h.providers[key]; ok {
		return provider
	}
	return nil
}

// 实例化一个服务
func (h *SimpleContainer) newInstance(provider ServiceProvider, params []any) (any, error) {
	if err := provider.Boot(h); err != nil {
		return nil, err
	}
	if params == nil {
		params = provider.Params(h)
	}
	method := provider.Register(h)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, nil

}

// 真正的实例化一个服务
func (h *SimpleContainer) make(key string, params []any, forceNew bool) (any, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	provider := h.findServiceProvider(key)
	if provider == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	if forceNew {
		return h.newInstance(provider, params)
	}

	// 不需要强制重新实例化，如果容器中已经实例化了，那么就直接使用容器中的实例
	if ins, ok := h.instances[key]; ok {
		return ins, nil
	}
	// 容器中没有实例化，那么就实例化一个
	instance, err := h.newInstance(provider, nil)
	if err != nil {
		return nil, err
	}
	h.instances[key] = instance
	return instance, nil
}

// Bind 将服务容器和关键字做了绑定
func (h *SimpleContainer) Bind(provider ServiceProvider) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	key := provider.Name()

	h.providers[key] = provider

	if provider.IsDefer() == false {
		// 因为Boot方法中可能会调用Make方法,
		// 而Make方法中又会加锁,所以这里先释放锁,
		// 等Boot方法执行完毕后再加锁
		h.lock.Unlock() // 释放锁，防止死锁
		if err := provider.Boot(h); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(h)
		method := provider.Register(h)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		h.instances[key] = instance
		// 重新加锁
		h.lock.Lock()
	}
	return nil
}

func (h *SimpleContainer) IsBind(key string) bool {
	return h.findServiceProvider(key) != nil
}

func (h *SimpleContainer) Make(key string) (any, error) {
	return h.make(key, nil, false)
}

func (h *SimpleContainer) MustMake(key string) any {
	serv, err := h.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (h *SimpleContainer) MakeNew(key string, params []any) (any, error) {
	return h.make(key, params, true)
}
