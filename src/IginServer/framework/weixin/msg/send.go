package msg

import (
	"fmt"
	"time"
)

type WeixinXML struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   string `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
	PicUrl       string `xml:"PicUrl"`
}

func SendTXT(v WeixinXML, Content string) string {
	return fmt.Sprintf(text, v.FromUserName, v.ToUserName, time.Now().Unix(), "text", Content)
}
