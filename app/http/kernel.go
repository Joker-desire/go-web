/**
 * @Author: yy
 * @Description:
 * @File:  kernel
 * @Version: 1.0.0
 * @Date: 2023/06/08 12:48
 */

package http

import "github.com/Joker-desire/go-web/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	Routes(r)
	return r, nil
}
