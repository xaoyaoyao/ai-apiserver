/**
 * Package repo
 * @file      : router_user.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 13:01
 **/

package router

import "github.com/volcengine/skd/internal/common/endpoint"

type UserParams struct {
	ID   string `path:"id"`   // 绑定到 {id}
	Page int    `path:"page"` // 绑定到 {page}
	//Username string `path:"username"` // 绑定到 {username} path路径不允许为空的
}

func ProcessUserPosts(ctx *Context) {
	var params UserParams
	if err := BindPathParams(ctx.Request, endpoint.API_STSTEM_PATH+endpoint.USERS_POSTS_PATH, &params); err != nil {
		HandleError(ctx.Writer, err)
		return
	}
	data := make(map[string]interface{})
	data["id"] = params.ID
	data["page"] = params.Page
	data["username"] = "demo"
	ctx.makeOK(data)
}
