/**
 * Package repo
 * @file      : main_test.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/23 10:14
 **/

package main

import (
	"fmt"
	req2 "github.com/volcengine/skd/internal/common/req"
	"github.com/volcengine/skd/internal/common/resp"
	"github.com/volcengine/skd/internal/config"
	"github.com/volcengine/skd/internal/handler"
	"os"
	"strconv"
	"testing"
)

func init() {
	_ = config.Init()
}

func TestSyncPush(t *testing.T) {
	var initImages []req2.InitImage
	initImage := req2.InitImage{
		Url: "https://static.vobase.com/images/20250420/2981745319285_.pic.jpg",
	}
	initImages = append(initImages, initImage)

	syncPushReq := req2.SyncPushReq{
		SyncTimeout: 30,
		TaskType:    "mtlab",
		InitImages:  initImages,
		Task:        "v1/sod",
	}

	syncPushResp, err := handler.SyncPush(syncPushReq)
	fmt.Println(err)
	fmt.Println(syncPushResp)
}

func TestOverResolutionV2(t *testing.T) {
	imageUrls := []string{"https://static.vobase.com/images/20250420/1411745296603_.pic.jpg"}
	processRequest := req2.ProcessReq{
		ReqKey:    "lens_vida_nnsr",
		ImageUrls: &imageUrls,
	}

	processResponse, _, _ := handler.OverResolutionV2(processRequest)
	Write(processResponse)

}

func TestEntitySegment(t *testing.T) {
	imageUrls := []string{"https://static.vobase.com/images/20250420/1411745296603_.pic.jpg"}
	maxEntity := 100
	returnFormat := 4
	refineMask := 1
	processRequest := req2.EntitySegmentReq{
		ReqKey:       "entity_seg",
		ImageUrls:    &imageUrls,
		MaxEntity:    &maxEntity,
		ReturnFormat: &returnFormat,
		RefineMask:   &refineMask,
	}

	processResponse, _, _ := handler.EntitySegment(processRequest)

	Write(processResponse)
}

func TestCVProcess(t *testing.T) {
	imageUrls := []string{"https://static.vobase.com/images/20250420/1411745296603_.pic.jpg"}
	//subReqKey := "img2img_comic_style_future"
	processRequest := req2.ProcessReq{
		ReqKey: "img2img_ceramics_style_usage",
		//SubReqKey: &subReqKey,
		ImageUrls: &imageUrls,
	}

	processResponse, _, _ := handler.CVProcess(processRequest)

	Write(processResponse)

}

func Write(processResponse resp.ProcessResponse) {
	//binaryDataBase64 := processResponse.Data.BinaryDataBase64
	//rt := (*binaryDataBase64)[0]

	//rt, _ := json.Marshal(processResponse)
	//
	//file, err := os.Create("response.txt")
	//if err != nil {
	//	panic("创建文件失败: " + err.Error())
	//}
	//defer file.Close() // 确保关闭文件

	//_, err = file.Write([]byte(rt))
	//if err != nil {
	//	panic("写入文件失败: " + err.Error())
	//}

	binaryDataBase64 := processResponse.Data.BinaryDataBase64
	if binaryDataBase64 != nil && len(*binaryDataBase64) > 0 {
		for i := 0; i < len(*binaryDataBase64); i++ {
			base64 := (*binaryDataBase64)[i]
			filename := "html_" + strconv.Itoa(i) + ".html"
			file, err := os.Create(filename)
			htmlStr := "<html><body>"
			htmlStr += "<img src=\"data:image/jpeg;base64," + base64 + "\" />"
			htmlStr += "</body></html>"
			_, err = file.Write([]byte(htmlStr))
			if err != nil {

			}
			defer file.Close()
		}
	}
}
