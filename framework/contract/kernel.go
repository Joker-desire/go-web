/**
 * @Author: yy
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:42
 */

package contract

import "net/http"

// KernelKey 提供kernel服务凭证
const KernelKey = "go-web:kernel"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}
