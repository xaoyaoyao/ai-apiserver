/**
 * Package repo
 * @file      : entity_segment.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 19:20
 **/

package req

type EntitySegmentReq struct {
	ReqKey           string    `json:"req_key"`
	BinaryDataBase64 *[]string `json:"binary_data_base64,omitempty"`
	ImageUrls        *[]string `json:"image_urls,omitempty"`
	MaxEntity        *int      `json:"max_entity,omitempty"`
	ReturnFormat     *int      `json:"return_format,omitempty"`
	RefineMask       *int      `json:"refine_mask,omitempty"`
}
