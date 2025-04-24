/**
 * Package repo
 * @file      : handler_meitu.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 11:20
 **/

package handler

import (
	"github.com/volcengine/skd/internal/meitu"
	"github.com/volcengine/skd/internal/req"
	"github.com/volcengine/skd/internal/resp"
)

func SyncPush(syncPushReq req.SyncPushReq) (*resp.SyncPushResp, error) {
	return meitu.SyncPush(syncPushReq)
}
