package auth

import (
	"net/http"
	// "IginServer/lib/md5"
	// "IginServer/lib/Imartini"
	// "IginServer/lib/session"
	"IginServer/lib/mygin"
	"IginServer/web/model/tranca/admin/auth"
)

func Auth(c *mygin.IContext) {
	i := auth.Auth(c.GetAllSession())
	if i == 1 {
		return
	} else {
		http.Redirect(c.Writer, c.Request, "/admin/login?msg=登陆已超时&url="+c.Request.URL.Path, http.StatusFound)
		return
	}
}
