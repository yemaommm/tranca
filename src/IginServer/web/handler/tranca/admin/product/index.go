package product

import (
	"IginServer/lib/mygin"
	// "log"
	// "github.com/go-martini/martini"
	// "fmt"
	"net/http"
	// "server/conf"
	// "IginServer/lib/Imartini"
	"IginServer/framework/API/R"
	"IginServer/lib/cookie"
	"IginServer/lib/md5"
	"IginServer/lib/other"
	// "IginServer/lib/upload"
	"IginServer/web/model/tranca/admin/product"
	"strconv"
	"time"
)

func index(c *mygin.IContext) {
	// info := c.GetSession("admin_info").(map[string]string)

	var page, size int
	page, _ = strconv.Atoi(c.Request.FormValue("page"))
	if page <= 0 {
		page = 1
	}
	size = 10

	search := map[string]string{
		"name": cookie.Get(c.Request, "product:searchname"),
	}
	// fmt.Println(search)
	products, count := product.ProductList((page-1)*size, size, search)

	c.HTML("tranca/admin/product/list.html", map[string]interface{}{
		"products": products,
		"total":    count,
		"size":     size,
	})
}

func AddProduct(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	// path, _ := upload.ImgSave(c.Request, "imgurl", 720, 720, info["id"].(string))

	param := map[string]interface{}{
		"aid":        info["id"],
		"product_id": md5.GetUUID().Hex(),
		"name":       c.Request.FormValue("name"),
		"price":      other.ToFloat64(c.Request.FormValue("price")),
		"remark":     c.Request.FormValue("remark"),
		"body":       c.Request.FormValue("body"),
		"createtime": int(time.Now().Unix()),
	}
	product.AddProduct(param)
	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))

}

func DelProduct(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return goods.Delactivity(params["id"], info["id"].(string))
	if product.DelProduct(c.Param("product_id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func ProductInfo(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	product := product.ProductInfo(info["id"], c.Param("product_id"))

	if len(product) > 0 {
		R.Write(c.Writer, map[string]interface{}{"product": product[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func UpdateProduct(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	param := map[string]interface{}{
		"aid":    info["id"],
		"name":   c.Request.FormValue("name"),
		"price":  other.ToFloat64(c.Request.FormValue("price")),
		"remark": c.Request.FormValue("remark"),
		"body":   c.Request.FormValue("body"),
	}
	product.UpdateProduct(info["id"], c.Param("product_id"), param)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}
