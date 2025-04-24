/**
 * Package repo
 * @file      : path_parameter_test.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/24 13:45
 **/

package router

import (
	"fmt"
	"testing"
)

func TestCompilePattern(t *testing.T) {
	cleanPath := "/api/system/v1/users/1232/posts/2323"
	pattern := "/api/system/v1/users/{id}/posts/{page}"
	re, paramNames, err := compilePattern(pattern)
	fmt.Println(err)
	fmt.Println(re)
	fmt.Println(paramNames)

	matches := re.FindStringSubmatch(cleanPath)
	fmt.Printf("Matches: %v\n", matches)
	//fmt.Println(matches[0])
	//fmt.Println(matches[1])
	//fmt.Println(matches[2])
}
