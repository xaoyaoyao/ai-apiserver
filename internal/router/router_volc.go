/**
 * Package repo
 * @file      : router_volc.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 11:33
 **/

package router

import (
	"github.com/volcengine/skd/internal/action"
	"github.com/volcengine/skd/internal/handler"
	"github.com/volcengine/skd/internal/req"
	"github.com/volcengine/skd/internal/resp"
	"log"
	"net/http"
)

func ProcessVolcEngine(ctx *Context) {
	actionMethod, _ := action.ToVolcEngineAction(ctx.Query("action"))

	if actionMethod == action.EntitySegment {
		var entitySegment req.EntitySegmentReq
		err := ctx.BindJSON(&entitySegment)
		if err != nil {
			log.Printf("[%s] %s %s| Error\n", ctx.Request.Method, ctx.Request.URL.Path, err.Error())
			ctx.makeError(http.StatusInternalServerError, nil)
			return
		}
		processResponse, code, err := handler.Process(actionMethod, &entitySegment, nil)
		makeProcessVolcEngine(ctx, processResponse, code, err)
		return
	}

	var processRequest req.ProcessReq
	err := ctx.BindJSON(&processRequest)
	if err != nil {
		log.Printf("[%s] %s %s| Error\n", ctx.Request.Method, ctx.Request.URL.Path, err.Error())
		ctx.makeError(http.StatusInternalServerError, nil)
		return
	}
	processResponse, code, err := handler.Process(actionMethod, nil, &processRequest)
	makeProcessVolcEngine(ctx, processResponse, code, err)
}

func makeProcessVolcEngine(ctx *Context, processResponse *resp.ProcessResponse, code int, err error) {
	if err != nil {
		log.Printf("[%s] %s %s| Error\n", ctx.Request.Method, ctx.Request.URL.Path, err.Error())
		ctx.makeError(code, nil)
		return
	}
	ctx.makeOK(processResponse)
}
