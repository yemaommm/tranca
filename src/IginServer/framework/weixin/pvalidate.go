package weixin

import (
	// "github.com/martini-contrib/render"
	// "log"
	// "github.com/go-martini/martini"
	// "fmt"
	"encoding/xml"
	"io/ioutil"
	// "net/http"
	// "IginServer/conf"
	"IginServer/framework/weixin/msg"
	"IginServer/lib/Imartini"
	"github.com/gin-gonic/gin"
	// "time"
)

func pvalidate(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	Imartini.MyLog.Other("weixin.", "%s %s\n%s\n", c.Request.RequestURI, c.Request.Method, body)

	var v msg.WeixinXML
	xml.Unmarshal(body, &v)

	c.String(200, msg.SendTXT(v, v.Content))
}
