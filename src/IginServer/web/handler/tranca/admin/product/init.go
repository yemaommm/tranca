package product

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
	r.GET("/productlist", mygin.Handler(index))
	r.POST("/productlist", mygin.Handler(AddProduct))
	r.GET("/product/del/:product_id", mygin.Handler(DelProduct))
	r.GET("/product/info/:product_id", mygin.Handler(ProductInfo))
	r.POST("/product/update/:product_id", mygin.Handler(UpdateProduct))
}
