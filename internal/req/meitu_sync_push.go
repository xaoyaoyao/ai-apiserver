/**
 * Package repo
 * @file      : meitu_sync_push.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 10:25
 **/

package req

type SyncPushReq struct {
	SyncTimeout int          `json:"sync_timeout"`
	TaskType    string       `json:"task_type"`
	InitImages  []InitImage  `json:"init_images"`
	Task        string       `json:"task"`
	Params      *MeituParams `json:"params"`
}

type InitImage struct {
	Url string `json:"url"`
}

type MeituParams struct {
	Parameter map[string]interface{} `json:"rsp_media_type"`
}
