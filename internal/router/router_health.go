/**
 * Package repo
 * @file      : router_health.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 14:27
 **/

package router

func HealthCheck(ctx *Context) {
	ctx.makeOK(nil)
}
