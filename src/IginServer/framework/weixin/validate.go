package weixin

import (
	// "github.com/martini-contrib/render"
	// "log"
	// "github.com/go-martini/martini"
	"crypto/sha1"
	"fmt"
	// "net/http"
	"sort"
	"strings"
	// "IginServer/conf"
	// "IginServer/lib/Imartini"
	"github.com/gin-gonic/gin"
)

var (
	token = "negocat"
)

func validate(c *gin.Context) {
	signature := c.Request.FormValue("signature")
	timestamp := c.Request.FormValue("timestamp")
	nonce := c.Request.FormValue("nonce")
	echostr := c.Request.FormValue("echostr")

	sign := []string{token, timestamp, nonce}
	sort.Strings(sign)

	b := sha1.Sum([]byte(strings.Join(sign, "")))

	if signature == fmt.Sprintf("%x", b) {
		c.String(200, echostr)
	} else {
		c.String(200, "")
	}
}
