package auth

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
	r := Imartini.M.Group("api/auth")
	r.POST("register", Register)
	r.POST("login", Login)
	r.POST("logout", Logout)
	r.POST("/", Auth)
}
