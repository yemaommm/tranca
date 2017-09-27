package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type bm struct {
	Data map[string]interface{}
}

func Map2Byte(info map[string]interface{}) []byte {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, bm{Data: info})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func Byte2Map(b []byte) map[string]interface{} {
	var info bm
	buf := bytes.NewBuffer(b)

	err := binary.Read(buf, binary.LittleEndian, &info)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return info.Data
	// Output: 3.141592653589793
}

func main() {
	b := Map2Byte(map[string]interface{}{"1": 123, "2": map[string]interface{}{"ddd": "www"}})
	fmt.Println(b)
	fmt.Println(Byte2Map(b))
}
