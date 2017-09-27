package msg

import (
	"fmt"
)

type ResultXml struct {
	Return_code string `xml:"return_code"`
	Return_msg  string `xml:"return_msg"`
	Appid       string `xml:"appid"`
	Mch_id      string `xml:"mch_id"`
	Nonce_str   string `xml:"nonce_str"`
	Sign        string `xml:"sign"`
	Result_code string `xml:"result_code"`
	Prepay_id   string `xml:"prepay_id"`
	Trade_type  string `xml:"trade_type"`
	Mweb_url    string `xml:"mweb_url"`
	Code_url    string `xml:"code_url"`
}

func MapToXml(data map[string]interface{}) (body string) {
	str := "<%s>%v</%s>"
	stmp := ""
	for i, j := range data {
		stmp += fmt.Sprintf(str, i, j, i)
	}
	body = fmt.Sprintf(str, "xml", stmp, "xml")
	return
}
