/**
 * Package repo
 * @file      : vc_process.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 13:47
 **/

package action

import "strings"

var (
	img2img_ghibli_style_usage       = CVProcessParameters{"网红日漫风", "img2img_ghibli_style_usage", ""}
	img2img_disney_3d_style_usage    = CVProcessParameters{"3D风", "img2img_disney_3d_style_usage", ""}
	img2img_real_mix_style_usage     = CVProcessParameters{"写实风", "img2img_real_mix_style_usage", ""}
	img2img_pastel_boys_style_usage  = CVProcessParameters{"天使风", "img2img_pastel_boys_style_usage", ""}
	img2img_cartoon_style_usage      = CVProcessParameters{"动漫风", "img2img_cartoon_style_usage", ""}
	img2img_makoto_style_usage       = CVProcessParameters{"日漫风", "img2img_makoto_style_usage", ""}
	img2img_rev_animated_style_usage = CVProcessParameters{"公主风", "img2img_rev_animated_style_usage", ""}
	img2img_blueline_style_usage     = CVProcessParameters{"梦幻风", "img2img_blueline_style_usage", ""}
	img2img_water_ink_style_usage    = CVProcessParameters{"水墨风", "img2img_water_ink_style_usage", ""}
	i2i_ai_create_monet_usage        = CVProcessParameters{"新莫奈花园", "i2i_ai_create_monet_usage", ""}
	img2img_water_paint_style_usage  = CVProcessParameters{"水彩风", "img2img_water_paint_style_usage", ""}
	img2img_comic_style_monet        = CVProcessParameters{"莫奈花园", "img2img_comic_style_usage", "img2img_comic_style_monet"}
	img2img_comic_style_marvel       = CVProcessParameters{"精致美漫", "img2img_comic_style_usage", "img2img_comic_style_marvel"}
	img2img_comic_style_future       = CVProcessParameters{"赛博机械", "img2img_comic_style_usage", "img2img_comic_style_future"}
	img2img_exquisite_style_usage    = CVProcessParameters{"精致韩漫", "img2img_exquisite_style_usage", ""}
	img2img_pretty_style_ink         = CVProcessParameters{"国风-水墨", "img2img_pretty_style_usage", "img2img_pretty_style_ink"}
	img2img_pretty_style_light       = CVProcessParameters{"浪漫光影", "img2img_pretty_style_usage", "img2img_pretty_style_light"}
	img2img_ceramics_style_usage     = CVProcessParameters{"陶瓷娃娃", "img2img_ceramics_style_usage", ""}
	img2img_chinese_style_usage      = CVProcessParameters{"中国红", "img2img_chinese_style_usage", ""}
	img2img_clay_style_3d            = CVProcessParameters{"丑萌粘土", "img2img_clay_style_usage", "img2img_clay_style_3d"}
	img2img_clay_style_bubble        = CVProcessParameters{"可爱玩偶", "img2img_clay_style_usage", "img2img_clay_style_bubble"}
	img2img_3d_style_era             = CVProcessParameters{"3D-游戏_Z时代", "img2img_3d_style_usage", "img2img_3d_style_era"}
	img2img_3d_style_movie           = CVProcessParameters{"动画电影", "img2img_3d_style_usage", "img2img_3d_style_movie"}
	img2img_3d_style_doll            = CVProcessParameters{"玩偶", "img2img_3d_style_usage", "img2img_3d_style_doll"}

	availableCVProcessParameters = []CVProcessParameters{img2img_ghibli_style_usage, img2img_disney_3d_style_usage,
		img2img_real_mix_style_usage, img2img_pastel_boys_style_usage, img2img_cartoon_style_usage, img2img_makoto_style_usage, img2img_rev_animated_style_usage,
		img2img_blueline_style_usage, img2img_water_ink_style_usage, i2i_ai_create_monet_usage, img2img_water_paint_style_usage,
		img2img_comic_style_monet, img2img_comic_style_marvel, img2img_comic_style_future, img2img_exquisite_style_usage, img2img_pretty_style_ink,
		img2img_pretty_style_light, img2img_ceramics_style_usage, img2img_chinese_style_usage, img2img_clay_style_3d, img2img_clay_style_bubble,
		img2img_3d_style_era, img2img_3d_style_movie, img2img_3d_style_doll,
	}
)

type CVProcessParameters struct {
	text      string
	reqKey    string
	subReqKey string
}

func (c CVProcessParameters) ReqKey() string {
	return c.reqKey
}

func (c CVProcessParameters) SubReqKey() string {
	return c.subReqKey
}

func (c CVProcessParameters) Text() string {
	return c.text
}

func ToCVProcess(reqKey string) (CVProcessParameters, error) {
	for _, v := range availableCVProcessParameters {
		if strings.EqualFold(v.ReqKey(), reqKey) {
			return v, nil
		}
	}
	return CVProcessParameters{
		reqKey:    img2img_ghibli_style_usage.ReqKey(),
		subReqKey: img2img_ghibli_style_usage.SubReqKey(),
		text:      img2img_ghibli_style_usage.Text(),
	}, nil
}
