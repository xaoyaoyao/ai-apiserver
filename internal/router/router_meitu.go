/**
 * Package repo
 * @file      : router_meitu.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 11:34
 **/

package router

import (
	"github.com/volcengine/skd/internal/common/req"
	"github.com/volcengine/skd/internal/handler"
	"log"
	"net/http"
)

func ProcessMeitu(ctx *Context) {
	//actionMethod, _ := action.ToMeituAction(ctx.Query("action"))

	var syncPushReq req.SyncPushReq
	err := ctx.BindJSON(&syncPushReq)
	if err != nil {
		log.Printf("[%s] %s %s| Error\n", ctx.Request.Method, ctx.Request.URL.Path, err.Error())
		ctx.makeError(http.StatusInternalServerError, nil)
		return
	}
	syncPushResp, err := handler.SyncPush(syncPushReq)
	if err != nil {
		log.Printf("[%s] %s %s| Error\n", ctx.Request.Method, ctx.Request.URL.Path, err.Error())
		ctx.makeError(http.StatusInternalServerError, nil)
		return
	}
	ctx.makeOK(syncPushResp)
}
