package weixin

import (
	// "encoding/json"
	// "fmt"
	// "github.com/go-martini/martini"
	// "github.com/martini-contrib/render"
	// "IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/session"
	"IginServer/web/model/tranca/admin/weixin"
	"net/http"
	// "strconv"
	"strings"
	// "time"
)

func weixinsetting(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	data := weixin.Weixininfo(info["id"])

	// mlog.Printf("%v", info["id"])

	c.HTML("tranca/admin/weixin/index.html", map[string]interface{}{
		"request": c.Request,
		"weixin":  data,
		// "total":   total,
	})
}

func updateweixin(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	p := make(map[string]interface{})
	// p["uid"] = info["id"].(string)
	p["appid"] = strings.TrimSpace(c.Request.FormValue("appid"))
	p["mchid"] = strings.TrimSpace(c.Request.FormValue("mchid"))
	p["key"] = strings.TrimSpace(c.Request.FormValue("key"))
	p["appsecret"] = strings.TrimSpace(c.Request.FormValue("appsecret"))

	weixin.UpdateWeixininfo(info["id"], p)

	weixin.SetWeixinConfig()
	c.Redirect(http.StatusFound, "/admin/weixinsetting")
}

func init() {
	weixin.SetWeixinConfig()
}
