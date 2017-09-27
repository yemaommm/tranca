package auth

import (
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/web"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin")
	r.GET("/login", mygin.Handler(Login))
	r.POST("/login", mygin.Handler(PLogin))
}
