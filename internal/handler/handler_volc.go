/**
 * Package repo
 * @file      : handler_volc.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 16:35
 **/

package handler

import (
	"encoding/json"
	"github.com/volcengine/skd/internal/common/action"
	req2 "github.com/volcengine/skd/internal/common/req"
	"github.com/volcengine/skd/internal/common/resp"
	"github.com/volcengine/skd/internal/common/util"
	"github.com/volcengine/skd/internal/config"
	"github.com/volcengine/skd/internal/service/volc"
	"net/url"
)

func Process(actionMethod action.VolcEngineAction, entitySegment *req2.EntitySegmentReq,
	processRequest *req2.ProcessReq) (*resp.ProcessResponse, int, error) {
	if actionMethod == action.EntitySegment {
		//entitySegment.ReqKey = "entity_seg"
		processResponse, code, err := EntitySegment(*entitySegment)
		return &processResponse, code, err
	}

	if actionMethod == action.OverResolutionV2 {
		//processRequest.ReqKey = "lens_vida_nnsr"
		processResponse, code, err := OverResolutionV2(*processRequest)
		return &processResponse, code, err
	}

	processResponse, code, err := CVProcess(*processRequest)
	return &processResponse, code, err
}

func CVProcess(req req2.ProcessReq) (resp.ProcessResponse, int, error) {
	return DoRequest(req, action.CVProcess, "POST")
}

func EntitySegment(req req2.EntitySegmentReq) (resp.ProcessResponse, int, error) {
	return DoRequest(req, action.EntitySegment, "POST")
}

func OverResolutionV2(req req2.ProcessReq) (resp.ProcessResponse, int, error) {
	return DoRequest(req, action.OverResolutionV2, "POST")
}

func DoRequest(req interface{}, action action.VolcEngineAction, method string) (resp.ProcessResponse, int, error) {
	resultMap, _ := util.StructToMap(req)
	reqBodyStr, _ := json.Marshal(resultMap)
	version := config.Get().VolcEngineVersion
	_, responseRaw, status, err := volc.DoRequest(action.Action(), version, method, url.Values{}, reqBodyStr)
	var processResponse resp.ProcessResponse
	err = json.Unmarshal(responseRaw, &processResponse)
	return processResponse, status, err
}
