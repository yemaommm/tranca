package admin

import (
	"IginServer/lib/mygin"
	// "log"
	// "github.com/go-martini/martini"
	// "net/http"
	// "fmt"
	// "server/conf"
	// "IginServer/lib/Imartini"
)

func index(c *mygin.IContext) {
	// , render.HTMLOptions{Layout: "shop/admin/base/base"}
	c.HTML("tranca/admin/index.html", map[string]interface{}{})
	// http.Redirect(res, req, Imartini.UrlFor("admin/toppage/setting"), http.StatusFound)
}
