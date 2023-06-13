/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:50
 */

package kernel

import (
	"github.com/Joker-desire/simple/framework/gin"
	"net/http"
)

type SimpleKernelService struct {
	engine *gin.Engine
}

func NewSimpleKernelService(params ...any) (any, error) {
	httpEngine := params[0].(*gin.Engine)
	return &SimpleKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *SimpleKernelService) HttpEngine() http.Handler {
	return s.engine
}
