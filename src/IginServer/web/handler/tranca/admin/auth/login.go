package auth

import (
	// "github.com/martini-contrib/render"
	// "log"
	"IginServer/lib/md5"
	"IginServer/lib/mygin"
	"net/http"
	// "IginServer/lib/session"
	"IginServer/web/model/tranca/admin/auth"
	// "fmt"
	// "IginServer/conf"
	// "IginServer/lib/Imartini"
)

func Login(c *mygin.IContext) {
	ret := map[string]interface{}{}

	ret["msg"] = c.Request.FormValue("msg")
	c.HTML("tranca/admin/login.html", ret)
}

func PLogin(c *mygin.IContext) {
	url := c.Request.FormValue("url")
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	if username == "" || password == "" {
		c.Redirect(http.StatusFound, "/admin/login?msg=用户名密码不能为空")
		return
	}

	i, dauth := auth.Login(username, md5.Md5(password))
	if i == 1 {
		c.SetSession("admin_user", username)
		c.SetSession("admin_passwd", md5.Md5(password))
		c.SetSession("admin_info", dauth)
		if url != "" {
			c.Redirect(http.StatusFound, url)
		} else {
			c.Redirect(http.StatusFound, "/admin/")
		}
	} else {
		c.Redirect(http.StatusFound, "/admin/login?msg=用户名或密码错误")
	}
}
