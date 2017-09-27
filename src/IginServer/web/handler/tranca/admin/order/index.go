package order

import (
	"IginServer/lib/mygin"
	// "log"
	// "github.com/go-martini/martini"
	// "fmt"
	"IginServer/conf"
	"net/http"
	// "IginServer/lib/Imartini"
	"IginServer/framework/API/R"
	"IginServer/lib/cookie"
	"IginServer/lib/md5"
	"IginServer/lib/other"
	"IginServer/lib/upload"
	"IginServer/web/model/tranca/admin/order"
	adminproduct "IginServer/web/model/tranca/admin/product"
	"IginServer/web/model/tranca/product"
	"strconv"
	// "time"
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
		"a.product_name":  cookie.Get(c.Request, "order:searchproductname"),
		"a.order_info_id": cookie.Get(c.Request, "order:searchorderid"),
		"b.username":      cookie.Get(c.Request, "order:searchusername"),
		"a.flow":          cookie.Get(c.Request, "order:searchtype"),
	}
	// fmt.Println(search)
	orders, count := order.OrderList((page-1)*size, size, search)
	products, _ := adminproduct.ProductList(0, 100, map[string]string{})

	c.HTML("tranca/admin/order/list.html", map[string]interface{}{
		"products": products,
		"orders":   orders,
		"total":    count,
		"size":     size,
	})
}

func DelOrder(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	//0：删除失败，1：删除成功
	// return goods.Delactivity(params["id"], info["id"].(string))
	if order.DelOrder(c.Param("order_id"), info["id"]) == 0 {
		c.String(200, "0")
	} else {
		c.String(200, "1")
	}

}

func OrderInfo(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	order := order.OrderInfo(info["id"], c.Param("order_id"))

	if len(order) > 0 {
		R.Write(c.Writer, map[string]interface{}{"order": order[1]})
	} else {
		R.Write(c.Writer, map[string]interface{}{})
	}
}

func UpdateOrder(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	business_licence, _ := upload.ImgSave(c.Request, "business_licence", 720, 720, info["id"])
	tax_registration_certificate, _ := upload.ImgSave(c.Request, "tax_registration_certificate", 720, 720, info["id"])
	Organization_Code_Certificate, _ := upload.ImgSave(c.Request, "Organization_Code_Certificate", 720, 720, info["id"])

	param := map[string]interface{}{
		"aid":     info["id"],
		"flow":    c.Request.FormValue("flow"),
		"phone":   c.Request.FormValue("phone"),
		"price":   other.ToFloat64(c.Request.FormValue("price")),
		"mail":    c.Request.FormValue("mail"),
		"company": c.Request.FormValue("company"),
	}
	if business_licence != "" {
		param["business_licence"] = conf.GetString("config", "IMAGE_HOST") + business_licence
	}
	if tax_registration_certificate != "" {
		param["tax_registration_certificate"] = conf.GetString("config", "IMAGE_HOST") + tax_registration_certificate
	}
	if Organization_Code_Certificate != "" {
		param["Organization_Code_Certificate"] = conf.GetString("config", "IMAGE_HOST") + Organization_Code_Certificate
	}
	order.UpdateOrder(info["id"], c.Param("order_id"), param)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}

func AddOrder(c *mygin.IContext) {
	info := c.GetSession("admin_info").(map[string]string)

	product_id := c.PostForm("product_id")
	company := c.PostForm("company")
	phone := c.PostForm("phone")
	mail := c.PostForm("mail")

	business_licence, _ := upload.ImgSave(c.Request, "business_licence",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), "0")
	tax_registration_certificate, _ := upload.ImgSave(c.Request, "tax_registration_certificate",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), "0")
	Organization_Code_Certificate, _ := upload.ImgSave(c.Request, "Organization_Code_Certificate",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), "0")

	if business_licence != "" {
		business_licence = conf.GetString("config", "IMAGE_HOST") + business_licence
	}
	if tax_registration_certificate != "" {
		tax_registration_certificate = conf.GetString("config", "IMAGE_HOST") + tax_registration_certificate
	}
	if Organization_Code_Certificate != "" {
		Organization_Code_Certificate = conf.GetString("config", "IMAGE_HOST") + Organization_Code_Certificate
	}

	productinfo := product.ProductFind(product_id)
	if len(productinfo) <= 0 {
		c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
		return
	}
	product_name := productinfo["name"]
	price := productinfo["price"]

	data := map[string]interface{}{
		"order_info_id":                 md5.GetUUID().Hex(),
		"uid":                           "0",
		"aid":                           info["id"],
		"product_id":                    product_id,
		"product_name":                  product_name,
		"price":                         price,
		"company":                       company,
		"phone":                         phone,
		"mail":                          mail,
		"business_licence":              business_licence,
		"tax_registration_certificate":  tax_registration_certificate,
		"Organization_Code_Certificate": Organization_Code_Certificate,
	}
	product.CreateOrderInfo(data)

	c.Redirect(http.StatusFound, c.Request.Header.Get("Referer"))
}
