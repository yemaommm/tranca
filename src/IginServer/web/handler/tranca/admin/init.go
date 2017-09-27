package admin

import (
	"IginServer/framework/ueditor"
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/web/handler/tranca/admin/auth"
	// "IginServer/web"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin", mygin.Handler(auth.Auth))
	r.GET("/", mygin.Handler(index))
	r.GET("/ueditor.config.js", ueditor.Ueditor_config)
	r.GET("/controller.php", ueditor.Controller)
	r.POST("/controller.php", ueditor.Controller)
}

// {"", "GET", "/ueditor.config.js", ueditor.Ueditor_config},
// {"", "GET", "/controller.php", ueditor.Controller},
// {"", "POST", "/controller.php", ueditor.Controller},
// {"admin/index", "GET", "", auth.Auth, index},
// {"admin/login", "GET", "/login", auth.Login},
// {"admin/plogin", "POST", "/login", auth.PLogin},
// {"admin/register", "GET", "/register", auth.Register},
// {"admin/pregister", "POST", "/register", auth.PRegister},
