/**
 * Package repo
 * @file      : tools_test.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 18:10
 **/

package util

import (
	"fmt"
	"github.com/volcengine/skd/internal/req"
	"testing"
)

type C struct {
	Depth int
}
type B struct {
	Toy    string
	Nested *C
}
type A struct {
	Name  string
	Age   int
	Child *B
}

func TestStructToMap(t *testing.T) {
	imageUrls := []string{"https://static.vobase.com/images/20250420/1411745296603_.pic.jpg"}
	processRequest := req.ProcessReq{
		ReqKey:    "img2img_ghibli_style_usage",
		ImageUrls: &imageUrls,
	}
	resultMap, _ := StructToMap(processRequest)
	fmt.Println(resultMap)

	position := 0
	subReqKey := "img2img_ghibli_style_usage"
	processRequest = req.ProcessReq{
		ReqKey:    "img2img_ghibli_style_usage",
		SubReqKey: &subReqKey,
		ImageUrls: &imageUrls,
		LogoInfo: &req.LogoInfoReq{
			Position: &position,
		},
	}
	resultMap, _ = StructToMap(processRequest)
	fmt.Println(resultMap)

	// 示例 1：Child 字段为空值
	a1 := A{Name: "Alice", Age: 0, Child: &B{Toy: "", Nested: &C{Depth: 0}}}
	rt, _ := StructToMap(a1)
	fmt.Printf("Case 1:\n%#v\n\n", rt) // 输出: map[string]interface {}{"Name":"Alice"}

	// 示例 2：多层嵌套非空
	a2 := A{Name: "Bob", Age: 30, Child: &B{Toy: "Car", Nested: &C{Depth: 5}}}
	rt, _ = StructToMap(a2)
	fmt.Printf("Case 2:\n%#v\n\n", rt) // 输出包含嵌套 Map

	// 示例 3：指针字段部分为空
	a3 := A{Name: "Charlie", Child: &B{Toy: "Ball", Nested: nil}}
	rt, _ = StructToMap(a3)
	fmt.Printf("Case 3:\n%#v\n\n", rt) // 输出: Child 包含 "Toy":"Ball"

}
