package other

import (
	// "bytes"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "os"
	"regexp"
	"strings"
)

func HtmlDecode(body string) string {
	body = strings.Replace(body, "&quot;", "\"", -1)
	body = strings.Replace(body, "&#39;", "'", -1)
	body = strings.Replace(body, "&nbsp;", " ", -1)
	body = strings.Replace(body, "&gt;", ">", -1)
	body = strings.Replace(body, "&lt;", "<", -1)
	body = strings.Replace(body, "&amp;", "&", -1)
	return body
}

func HtmlEncode(body string) string {
	body = strings.Replace(body, "&", "&amp;", -1)
	body = strings.Replace(body, "<", "&lt;", -1)
	body = strings.Replace(body, ">", "&gt;", -1)
	body = strings.Replace(body, " ", "&nbsp;", -1)
	body = strings.Replace(body, "'", "&#39;", -1)
	body = strings.Replace(body, "\"", "&quot;", -1)
	return body
}

func IsPhone(mobileNum string) bool {
	regular := "^(13[0-9]|14[57]|15[0-35-9]|17[0-9]|18[0-9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
