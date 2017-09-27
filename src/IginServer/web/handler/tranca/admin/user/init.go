package user

import (
	// "IginServer/framework/ueditor"
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
	r.GET("/userlist", mygin.Handler(index))
	r.POST("/wapuser/updatepwd/:uid", mygin.Handler(UpdateWapUserPassword))
}
