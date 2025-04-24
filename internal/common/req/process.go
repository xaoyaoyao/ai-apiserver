/**
 * Package repo
 * @file      : process.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 13:39
 **/

package req

type ProcessReq struct {
	ReqKey           string       `json:"req_key"`
	SubReqKey        *string      `json:"sub_req_key"`
	BinaryDataBase64 *[]string    `json:"binary_data_base64,omitempty"`
	ImageUrls        *[]string    `json:"image_urls,omitempty"`
	ReturnUrl        *bool        `json:"return_url,omitempty"`
	LogoInfo         *LogoInfoReq `json:"logo_info,omitempty"`
}

type LogoInfoReq struct {
	Position        *int     `json:"position,omitempty"`
	Language        *int     `json:"language,omitempty"`
	Opacity         *float32 `json:"opacity,omitempty"`
	LogoTextContent *string  `json:"logo_text_content,omitempty"`
}
