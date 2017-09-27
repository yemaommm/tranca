package other

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	// "regexp"
	"strconv"
	"strings"
)

/* struct转map */
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func XML2Map(data []byte, v string) map[string]interface{} {
	ret := make(map[string]interface{})
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	// name := ""
	v = strings.ToUpper(v)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			if strings.ToUpper(token.Name.Local) != v {
				// name = token.Name.Local
				var text string
				decoder.DecodeElement(&text, &token)
				ret[token.Name.Local] = text
				// fmt.Println(token.Name.Local, text)
				// fmt.Println("start:", token.Name.Local)
			}
		// 处理元素结束（标签）
		case xml.EndElement:
			// if strings.ToUpper(token.Name.Local) != v {
			// name = token.Name.Local
			// fmt.Println("end:", token.Name.Local)
			// }
		// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			// if strings.ToUpper(name) != v {
			// 	text := string([]byte(token))
			// 	if name != "" && (ret[name] == "" || ret[name] == nil) {
			// 		// fmt.Println("value:", text)
			// 		// fmt.Println(len(string([]byte(token))))
			// 		ret[name] = text
			// 	}
			// }
		default:
			// fmt.Println(token)
		}
		// fmt.Println(name)
	}
	return ret
}

func Map2XML(data map[string]interface{}) string {
	ret := []string{"<xml>"}

	for i, j := range data {
		v := fmt.Sprintf("%v", j)
		_, err := strconv.ParseFloat(v, 64)
		if err != nil {
			ret = append(ret, fmt.Sprintf("<%s><![CDATA[%s]]></%s>", i, v, i))
		} else {
			ret = append(ret, fmt.Sprintf("<%s>%s</%s>", i, v, i))
		}
	}
	ret = append(ret, "</xml>")

	return strings.Join(ret, "\n")
}

func Map2Json(data interface{}) string {
	ret, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return string(ret)
}

func Json2Map(str string) map[string]interface{} {
	var info map[string]interface{}
	json.Unmarshal([]byte(str), &info)

	return info
}
