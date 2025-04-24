/**
 * Package repo
 * @file      : http.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 10:16
 **/

package meitu

import (
	"encoding/json"
	"github.com/volcengine/skd/internal/config"
	"github.com/volcengine/skd/internal/req"
	"github.com/volcengine/skd/internal/resp"
	"github.com/volcengine/skd/internal/util"
	"io"
	"net/http"
)

func SyncPush(syncPushReq req.SyncPushReq) (*resp.SyncPushResp, error) {
	sign := NewSigner(config.Get().MeituApiKey, config.Get().MeituSecretKey)
	url := config.Get().MeituSyncPushUrl
	method := "POST"
	headers := make(http.Header)
	headers.Set(HeaderHost, "openapi.meitu.com")
	headers.Set("Content-Type", "application/json")

	resultMap, _ := util.StructToMap(syncPushReq)
	reqBody, _ := json.Marshal(resultMap)

	request, err := sign.Sign(url, method, headers, string(reqBody))
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var syncPushResp resp.SyncPushResp
	err = json.Unmarshal(body, &syncPushResp)
	return &syncPushResp, nil
}
