/**
 * Package repo
 * @file      : response.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 18:09
 **/

package util

import (
	"fmt"
	"reflect"
	"strings"
)

func StructToMap(input interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(input)

	// 解引用指针，直到获取实际的值
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return map[string]interface{}{}, nil
		}
		v = v.Elem()
	}

	// 仅处理结构体
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or pointer to struct")
	}

	result := make(map[string]interface{})
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// 解析 JSON 标签
		tag := field.Tag.Get("json")
		key := strings.Split(tag, ",")[0]
		if key == "" || key == "-" {
			continue // 忽略无标签或明确忽略的字段
		}

		var value interface{}

		// 处理指针类型字段
		if field.Type.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue // 指针为 nil 则跳过
			}
			elemValue := fieldValue.Elem()

			// 递归处理嵌套结构体
			if elemValue.Kind() == reflect.Struct {
				subMap, err := StructToMap(elemValue.Interface())
				if err != nil {
					return nil, err
				}
				if len(subMap) > 0 {
					value = subMap
				} else {
					continue // 子结构体为空，不添加
				}
			} else {
				value = elemValue.Interface()
			}

		} else {
			// 处理值类型字段
			if fieldValue.IsZero() {
				continue // 零值则跳过
			}

			// 递归处理嵌套结构体
			if fieldValue.Kind() == reflect.Struct {
				subMap, err := StructToMap(fieldValue.Interface())
				if err != nil {
					return nil, err
				}
				if len(subMap) > 0 {
					value = subMap
				} else {
					continue // 子结构体为空，不添加
				}
			} else {
				value = fieldValue.Interface()
			}
		}

		result[key] = value
	}

	return result, nil
}
