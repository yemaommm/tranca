package user

import (
	"IginServer/lib/mygin"
	// "log"
	// "github.com/go-martini/martini"
	"net/http"
	// "fmt"
	// "server/conf"
	// "IginServer/lib/Imartini"
	"IginServer/lib/cookie"
	"IginServer/web/model/tranca/admin/user"
	"strconv"
)

func index(c *mygin.IContext) {
	// , render.HTMLOptions{Layout: "shop/admin/base/base"}
	// c.HTML("tranca/admin/user/list.html", map[string]interface{}{})
	// http.Redirect(res, req, Imartini.UrlFor("admin/toppage/setting"), http.StatusFound)

	// sess := session.GetAll(req)
	// info := c.GetSession("admin_info").(map[string]string)

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	if page <= 0 {
		page = 1
	}
	size = 10

	search := map[string]string{
		"username": cookie.Get(c.Request, "user:searchusername"),
	}
	// fmt.Println(search)
	userlist, count := user.WapUserList((page-1)*size, size, search)

	c.HTML("tranca/admin/user/list.html", map[string]interface{}{
		"userlist": userlist,
		"total":    count,
		"size":     size,
	})
}

func UpdateWapUserPassword(c *mygin.IContext) {
	// , render.HTMLOptions{Layout: "shop/admin/base/base"}
	// c.HTML("tranca/admin/user/list.html", map[string]interface{}{})
	// http.Redirect(res, req, Imartini.UrlFor("admin/toppage/setting"), http.StatusFound)

	// sess := session.GetAll(req)
	// info := c.GetSession("admin_info").(map[string]string)

	uid := c.Param("uid")
	password := c.PostForm("password")
	// fmt.Println(search)
	user.UpdateWapUserPassword(uid, password)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}
