/**
 * Package repo
 * @file      : meitu.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 11:09
 **/

package action

import (
	"fmt"
	"strings"
)

var (
	ErrUnknownMeituAction = fmt.Errorf("unknown Meitu VolcEngineAction")
)

var (
	SyncPush              = MeituAction{"SyncPush"}
	availableMeituActions = []MeituAction{SyncPush}
)

type MeituAction struct {
	action string
}

func (a MeituAction) Action() string {
	return a.action
}

func ToMeituAction(action string) (MeituAction, error) {
	for _, v := range availableMeituActions {
		if strings.EqualFold(v.Action(), action) {
			return v, nil
		}
	}
	return MeituAction{
		action: SyncPush.Action(),
	}, nil
}
