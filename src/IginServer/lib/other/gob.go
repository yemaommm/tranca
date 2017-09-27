package other

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func ToByte(info interface{}) []byte {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)
	err := enc.Encode([]interface{}{info})
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return []byte{}
	}
	return buf.Bytes()
}

func To(b []byte) interface{} {
	var info []interface{}
	buf := bytes.NewBuffer(b)

	dec := gob.NewDecoder(buf)
	err := dec.Decode(&info)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
		return nil
	}
	return info[0]
	// Output: 3.141592653589793
}

func init() {
	gob.Register(map[string]string{})
	gob.Register([]map[string]string{})
	gob.Register(map[string]interface{}{})
	gob.Register(map[int]map[string]string{})

	// b := ToByte(map[string]interface{}{"1": 123, "2": map[string]interface{}{"ddd": "www"}})
	// fmt.Println(b)
	// fmt.Println(To(b))
}
