package order

import (
	// "IginServer/framework/ueditor"
	"IginServer/lib/Imartini"
	"IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/web/handler/tranca/admin/auth"
	// "encoding/json"
	// "fmt"
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "time"
)

func init() {
	r := Imartini.M.Group("admin", mygin.Handler(auth.Auth))
	r.GET("/orderlist", mygin.Handler(index))
	r.GET("/order/del/:order_id", mygin.Handler(DelOrder))
	r.GET("/order/info/:order_id", mygin.Handler(OrderInfo))
	r.POST("/order/update/:order_id", mygin.Handler(UpdateOrder))
	r.POST("/order/add", mygin.Handler(AddOrder))
}
