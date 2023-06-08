/**
 * @Author: yy
 * @Description:
 * @File:  route
 * @Version: 1.0.0
 * @Date: 2023/06/08 12:48
 */

package http

import (
	"github.com/Joker-desire/go-web/app/http/module/demo"
	"github.com/Joker-desire/go-web/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")

	_ = demo.Register(r)
}
