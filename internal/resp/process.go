/**
 * Package repo
 * @file      : process.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 13:39
 **/

package resp

type ProcessResponse struct {
	Code        int    `json:"code"`
	Data        Data   `json:"data"`
	Message     string `json:"message"`
	RequestId   string `json:"request_id"`
	Status      int    `json:"status"`
	TimeElapsed string `json:"time_elapsed"`
}

type Data struct {
	BinaryDataBase64  *[]string          `json:"binary_data_base64,omitempty"`
	ComfyuiCost       *int32             `json:"comfyui_cost,omitempty"`
	ImageOutput       *[]string          `json:"image_output,omitempty"`
	SaveImage935      *[]string          `json:"SaveImage_935,omitempty"`
	AlgorithmBaseResp *AlgorithmBaseResp `json:"algorithm_base_resp,omitempty"`
	ImageUrls         *[]string          `json:"image_urls,omitempty"`
	OriHeight         *[]int             `json:"ori_height,omitempty"`
	OriWidth          *[]int             `json:"ori_width,omitempty"`
	EntityNum         *[]int             `json:"entity_num,omitempty"`
	SegScore          *[]float64         `json:"seg_score,omitempty"`
}

type AlgorithmBaseResp struct {
	StatusCode    *int    `json:"status_code,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
}
