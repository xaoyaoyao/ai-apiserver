/**
 * Package repo
 * @file      : meitu_sync_psuh.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 10:41
 **/

package resp

type SyncPushResp struct {
	RequestId string           `json:"request_id"`
	TraceId   string           `json:"trace_id"`
	Code      int              `json:"code"`
	ErrorCode int              `json:"error_code"`
	Message   string           `json:"message"`
	Data      SyncPushRespData `json:"data"`
}

type SyncPushRespData struct {
	Status         int                `json:"status"`
	Result         SyncPushRespResult `json:"result"`
	Progress       int                `json:"progress"`
	PredictElapsed int                `json:"predict_elapsed"`
	CreateTime     int64              `json:"create_time"`
	TaskId         string             `json:"task_id"`
	CustomTaskId   string             `json:"custom_task_id"`
	TraceId        string             `json:"trace_id"`
	ClientInfo     string             `json:"client_info"`
	InitImages     interface{}        `json:"init_images"`
}

type SyncPushRespResult struct {
	Id            string                            `json:"id"`
	Urls          []string                          `json:"urls"`
	Parameters    *SyncPushRespParam                `json:"parameters"`
	Data          *SyncPushRespResultData           `json:"data"`
	Msg           string                            `json:"msg"`
	MsgId         string                            `json:"msg_id"`
	Images        []string                          `json:"images"`
	MediaInfoList []SyncPushRespResultMediaInfoList `json:"media_info_list"`
}

type SyncPushRespResultMediaInfoList struct {
	MediaData     string      `json:"media_data"`
	MediaExtra    interface{} `json:"media_extra"`
	MediaProfiles struct {
		MediaDataType string `json:"media_data_type"`
	} `json:"media_profiles"`
}

type SyncPushRespParam struct {
	Kind         int     `json:"Kind"`
	BottomX      string  `json:"bottom_x"`
	BottomY      string  `json:"bottom_y"`
	ExistSalient bool    `json:"exist_salient"`
	ProcessTime  float64 `json:"process_time"`
	PullTime     float64 `json:"pull_time"`
	RspMediaType string  `json:"rsp_media_type"`
	TopX         string  `json:"top_x"`
	TopY         string  `json:"top_y"`
	UseFe        bool    `json:"use_fe"`
	Version      string  `json:"version"`
}

type SyncPushRespResultData struct {
	Duration      *SyncPushRespResultDataDuration        `json:"duration"`
	ErrorCode     int                                    `json:"error_code"`
	ErrorMsg      string                                 `json:"error_msg"`
	MediaInfoList *[]SyncPushRespResultDataMediaInfoList `json:"media_info_list"`
	MsgId         string                                 `json:"msg_id"`
	Parameter     *SyncPushRespResultDataParameter       `json:"parameter"`
}

type SyncPushRespResultDataDuration struct {
	AlgProcessTime   int `json:"alg_process_time"`
	CreatedTimestamp int `json:"created_timestamp"`
	PullTimestamp    int `json:"pull_timestamp"`
	RepostTime       int `json:"repost_time"`
	UploadTime       int `json:"upload_time"`
	WaitingTime      int `json:"waiting_time"`
}

type SyncPushRespResultDataMediaInfoList struct {
	MediaData     string      `json:"media_data"`
	MediaExtra    interface{} `json:"media_extra"`
	MediaProfiles struct {
		MediaDataType string `json:"media_data_type"`
	} `json:"media_profiles"`
}

type SyncPushRespResultDataParameter struct {
	Kind         int     `json:"Kind"`
	BottomX      string  `json:"bottom_x"`
	BottomY      string  `json:"bottom_y"`
	ExistSalient bool    `json:"exist_salient"`
	ProcessTime  float64 `json:"process_time"`
	PullTime     float64 `json:"pull_time"`
	RspMediaType string  `json:"rsp_media_type"`
	TopX         string  `json:"top_x"`
	TopY         string  `json:"top_y"`
	UseFe        bool    `json:"use_fe"`
	Version      string  `json:"version"`
}
