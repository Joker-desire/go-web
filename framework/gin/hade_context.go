/**
 * @Author: yy
 * @Description:
 * @File:  hade_context
 * @Version: 1.0.0
 * @Date: 2023/06/07 10:54
 */

package gin

import "context"

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}
