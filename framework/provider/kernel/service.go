/**
 * @Author: yy
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2023/06/09 12:50
 */

package kernel

import (
	"github.com/Joker-desire/go-web/framework/gin"
	"net/http"
)

type HadeKernelService struct {
	engine *gin.Engine
}

func NewHadeKernelService(params ...any) (any, error) {
	httpEngine := params[0].(*gin.Engine)
	return &HadeKernelService{engine: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *HadeKernelService) HttpEngine() http.Handler {
	return s.engine
}
