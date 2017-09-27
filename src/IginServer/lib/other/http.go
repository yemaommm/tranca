package other

import (
	"bytes"
	"fmt"
	"io/ioutil"
	// "log"
	"compress/gzip"
	"net/http"
	// "os"
	"strconv"
)

func Httpget(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return []byte("")
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return []byte("")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func HttpCopyGet(req *http.Request, url string) []byte {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", url, nil)
	reqest.Header = req.Header
	response, _ := client.Do(reqest)

	if response.StatusCode != http.StatusOK {
		fmt.Println(response.StatusCode)
		return []byte("")
	}
	defer response.Body.Close()
	var body []byte
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		defer reader.Close()
		body, _ = ioutil.ReadAll(reader)
	default:
		bodyByte, _ := ioutil.ReadAll(response.Body)
		body = bodyByte
	}
	return body
}

func HttpPostBody(url string, msg []byte) []byte {
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(msg))
	if err != nil {
		fmt.Println(err)
		return []byte("")
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return []byte("")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func ToInt(v interface{}) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%v", v))
	return i
}

func ToFloat32(v interface{}) float32 {
	i, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 32)
	return float32(i)
}

func ToFloat64(v interface{}) float64 {
	i, _ := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	return i
}

func ToString(v interface{}) string {
	i := fmt.Sprintf("%v", v)
	return i
}

func Thow() {
	if err := recover(); err != nil {
		fmt.Println(err) //这里的err其实就是panic传入的内容，55
	}
}
