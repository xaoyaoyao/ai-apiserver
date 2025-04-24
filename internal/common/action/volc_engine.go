/**
 * Package repo
 * @file      : volc_engine.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 13:41
 **/

package action

import (
	"fmt"
	"strings"
)

var (
	ErrUnknownAction = fmt.Errorf("unknown VolcEngine Action")
)

var (
	CVProcess        = VolcEngineAction{"CVProcess"}
	EntitySegment    = VolcEngineAction{"EntitySegment"}
	OverResolutionV2 = VolcEngineAction{"OverResolutionV2"}
	availableActions = []VolcEngineAction{CVProcess, EntitySegment, OverResolutionV2}
)

type VolcEngineAction struct {
	action string
}

func (a VolcEngineAction) Action() string {
	return a.action
}

func ToVolcEngineAction(action string) (VolcEngineAction, error) {
	for _, v := range availableActions {
		if strings.EqualFold(v.Action(), action) {
			return v, nil
		}
	}
	return VolcEngineAction{
		action: CVProcess.Action(),
	}, nil
}
