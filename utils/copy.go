/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-22 10:18:28
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-03-24 21:51:02
 */

package utils


import (
	"encoding/json"
	"reflect"
	"strconv"
)

// Copy  b, _ := json.Marshal(v1) -> json.Unmarshal(b, &v2)
func Copy(v1 interface{}, v2 interface{}) error {
	b, err := json.Marshal(v1)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v2)
}

// CopyFromNames 通过字段名拷贝；暂时支持"string"转"uint","uint8","uint16"
func CopyFromNames(v1 interface{}, v2 interface{}, names []string) error {
	var m map[string]interface{}
	m = make(map[string]interface{})

	v1v := reflect.ValueOf(v1)
	v1t := reflect.TypeOf(v1)
	v2t := reflect.TypeOf(v2).Elem()

	for _, name := range names {
		v1tSF, _ := v1t.FieldByName(name)
		v2tSF, _ := v2t.FieldByName(name)
		v1vFBN := v1v.FieldByName(name)
		switch v1tSF.Type.Kind().String() {
			case "string":
				value := v1vFBN.String()
				switch v2tSF.Type.Kind().String() {
					case "uint","uint8","uint16":
						m[name], _ = strconv.Atoi(value)
					default:
						m[name] = value
				}
			case "uint","uint8","uint16","uint64":
				m[name] = v1vFBN.Uint()
		}
	}

	return Copy(m, v2)
}
