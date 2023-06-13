/**
 * @Author: yy
 * @Description:
 * @File:  distributed
 * @Version: 1.0.0
 * @Date: 2023/06/12 15:40
 */

package contract

import "time"

const DistributedKey = "simple:distributed"

// Distributed 定义分布式服务接口
type Distributed interface {
	// Select 分布式选择器, 所有节点对某个服务进行抢占，只选择其中一个节点
	// ServiceName 服务名字
	// appID 当前的AppID
	// holdTime 分布式选择器hold住的时间
	// 返回值
	// selectAppID 分布式选择器最终选择的App
	// err 异常才返回，如果没有被选择，不返回err
	Select(serviceName, appID string, holdTime time.Duration) (selectAppID string, err error)
}