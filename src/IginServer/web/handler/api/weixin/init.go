package weixin

import (
	"IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/web"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("api/weixin")
	r.POST("NativeNotify", NativeNotify)
	r.POST("NativePayNotify", NativePayNotify)
	r.GET("GetQRcodeUrl", GetQRcodeUrl)
	r.GET("GetRedirectUrl", GetRedirectUrl)
	r.GET("GetJsConfig", GetJsConfig)
}
