package product

import (
	"IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/web/handler/api/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("api/product")
	r.POST("findunpayorder", auth.AuthHandler, FindUnPayOrder)
	r.POST("createinfo", auth.AuthHandler, CreateInfo)
	r.GET("list", GetList)
}
