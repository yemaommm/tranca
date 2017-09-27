package product

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/conf"
	"IginServer/framework/API/R"
	// "IginServer/web/handler/api/sms"
	"IginServer/web/model/tranca/product"
	// "encoding/json"
	// w "IginServer/framework/weixin"
	"IginServer/lib/md5"
	// "IginServer/lib/other"
	// "IginServer/lib/redis"
	"IginServer/lib/upload"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	// "time"
)

/**
 * @api {post multipart/form-data} /api/product/createinfo 生成公司购买订单详情
 * @apiName 生成公司购买订单详情
 * @apiGroup 产品
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} token 登录返回的token
 * @apiParam {String} product_id 产品id
 * @apiParam {String} company 公司
 * @apiParam {String} phone 联系电话
 * @apiParam {String} mail 邮箱
 * @apiParam {String} business_licence 营业执照
 * @apiParam {String} tax_registration_certificate 税务登记证
 * @apiParam {String} Organization_Code_Certificate 组织机构代码证
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  成功，返回订单id
 * @apiSuccessExample Success-Response:
 *
 *
 * @apiSuccess (Reponse 202) {Number} status  202
 * @apiSuccess (Reponse 202) {String} error 创建失败
 *
 * @apiSuccess (Reponse 210) {Number} status  210
 * @apiSuccess (Reponse 210) {String} error 产品不存在
 *
 * @apiSuccess (Reponse 211) {Number} status  211
 * @apiSuccess (Reponse 211) {String} error 拥有未支付订单
 *
 * @apiSuccess (Reponse 400) {Number} status  400
 * @apiSuccess (Reponse 400) {String} error 登录验证失败
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/product/createinfo
 */
func CreateInfo(c *gin.Context) {
	g, _ := c.Get("userinfo")
	userinfo := g.(map[string]interface{})

	// if product.FindUnPayOrder(userinfo["uid"].(string)) != nil {
	// 	R.Error(c.Writer, 211, "拥有未支付订单")
	// 	return
	// }
	// username:=c.PostForm("username")
	// token:=c.PostForm("token")
	product_id := c.PostForm("product_id")
	company := c.PostForm("company")
	phone := c.PostForm("phone")
	mail := c.PostForm("mail")

	business_licence, _ := upload.ImgSave(c.Request, "business_licence",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), userinfo["uid"].(string))
	tax_registration_certificate, _ := upload.ImgSave(c.Request, "tax_registration_certificate",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), userinfo["uid"].(string))
	Organization_Code_Certificate, _ := upload.ImgSave(c.Request, "Organization_Code_Certificate",
		conf.GetInt("image", "SaveImgWidth"), conf.GetInt("image", "SaveImgHeight"), userinfo["uid"].(string))

	if company == "" {
		R.Api404(c.Writer)
		return
	} else if phone == "" {
		R.Api404(c.Writer)
		return
	} else if mail == "" {
		R.Api404(c.Writer)
		return
	} else if business_licence == "" {
		R.Api404(c.Writer)
		return
	} else if tax_registration_certificate == "" {
		R.Api404(c.Writer)
		return
	} else if Organization_Code_Certificate == "" {
		R.Api404(c.Writer)
		return
	}

	business_licence = conf.GetString("config", "IMAGE_HOST") + business_licence
	tax_registration_certificate = conf.GetString("config", "IMAGE_HOST") + tax_registration_certificate
	Organization_Code_Certificate = conf.GetString("config", "IMAGE_HOST") + Organization_Code_Certificate

	productinfo := product.ProductFind(product_id)
	if len(productinfo) <= 0 {
		R.Error(c.Writer, 210, "产品不存在")
		return
	}
	product_name := productinfo["name"]
	price := productinfo["price"]

	data := map[string]interface{}{
		"order_info_id":                 md5.GetUUID().Hex(),
		"uid":                           userinfo["uid"],
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
	i := product.CreateOrderInfo(data)

	if i <= 0 {
		R.Error(c.Writer, 202, "创建失败")
	} else {
		R.Success(c.Writer, data["order_info_id"])
	}
}

/**
 * @api {post} /api/product/findunpayorder 查询是否有未支付的订单产品
 * @apiName 查询是否有未支付的订单产品
 * @apiGroup 产品
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} username 用户名
 * @apiParam {String} token 登录返回的token
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  未有未支付订单，可以开始新的购买
 * @apiSuccessExample Success-Response:
 *
 *
 *
 * @apiSuccess (Reponse 211) {Number} status  211
 * @apiSuccess (Reponse 211) {String} error 拥有未支付订单，返回订单id
 *
 * @apiSuccess (Reponse 400) {Number} status  400
 * @apiSuccess (Reponse 400) {String} error 登录验证失败
 *
 * @apiSuccess (Reponse 404) {Number} status  404
 * @apiSuccess (Reponse 404) {String} error 缺少必要参数
 *
 * @apiSuccess (Reponse 403) {Number} status  403
 * @apiSuccess (Reponse 403) {String} error 参数不正确
 *
 * @apiError (Reponse 500) {Number} status 500
 * @apiError (Reponse 500) {String} error 系统错误
 *
 * @apiSampleRequest http://121.41.116.104:8003/api/product/findunpayorder
 */
func FindUnPayOrder(c *gin.Context) {
	g, _ := c.Get("userinfo")
	userinfo := g.(map[string]interface{})

	ret := product.FindUnPayOrder(userinfo["uid"].(string))

	if ret != nil {
		R.Error(c.Writer, 211, ret["order_info_id"])
		return
	}
	R.Success(c.Writer, "未有未支付订单，可以开始新的购买")
}
