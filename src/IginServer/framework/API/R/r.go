package R

import (
	"IginServer/lib/Imartini"
	"encoding/json"
	"fmt"
	"net/http"
	// "time"
)

func Write(res http.ResponseWriter, ret map[string]interface{}) error {
	b, err := json.Marshal(ret)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(200)
	res.Write(b)
	Imartini.AddLog(fmt.Sprintf("\napi:%s", b))
	return nil
}

func Api404(res http.ResponseWriter) error {
	return Write(res, map[string]interface{}{
		"status": 404,
		"error":  "缺少必要参数",
	})
}

func Success(res http.ResponseWriter, ret interface{}) error {
	return Write(res, map[string]interface{}{
		"status": 200,
		"data":   ret,
	})
}

func Error(res http.ResponseWriter, status int, e string) error {
	return Write(res, map[string]interface{}{
		"status": status,
		"error":  e,
	})
}
