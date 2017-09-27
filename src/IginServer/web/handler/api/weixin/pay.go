package weixin

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	"IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/web/model/tranca/product"
	// "encoding/json"
	w "IginServer/framework/weixin"
	wmsg "IginServer/framework/weixin/msg"
	// "IginServer/lib/md5"
	"IginServer/lib/other"
	"IginServer/lib/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	// "io/ioutil"
	// "net/http"
	"strings"
	"time"
)

/**
 * @api {get} /api/weixin/GetQRcodeUrl 获取订单扫码支付url
 * @apiName 获取订单扫码支付url
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} order_info_id 订单id
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回扫码的链接
 * @apiSuccessExample Success-Response:
 * {"data":"weixin://wxpay/bizpayurl?pr=qKcbbO7","status":200}
 *
 *
 *
 * @apiSuccess (Reponse 230) {Number} status  230
 * @apiSuccess (Reponse 230) {String} error 订单不存在
 *
 * @apiSuccess (Reponse 231) {Number} status  231
 * @apiSuccess (Reponse 231) {String} error 统一下单失败
 *
 * @apiSuccess (Reponse 410) {Number} status  410
 * @apiSuccess (Reponse 410) {String} error 错误参数
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
 * @apiSampleRequest http://121.41.116.104:8003/api/weixin/GetQRcodeUrl
 */
//获取订单扫码url
func GetQRcodeUrl(c *gin.Context) {
	order_info_id := c.Query("order_info_id")

	ret, code := ProductUnifiedOrder_pub("NATIVE", order_info_id, "", "")
	if code == 230 {
		R.Error(c.Writer, 230, "订单不存在")
		return
	} else if code == 231 {
		R.Error(c.Writer, 231, "统一下单失败")
		return
	}

	R.Success(c.Writer, ret.Code_url)
}

/**
 * @api {get} /api/weixin/GetRedirectUrl 获取订单跳转支付url
 * @apiName 获取订单跳转支付url
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} order_info_id 订单id
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回扫码的链接
 * @apiSuccessExample Success-Response:
 * {"data":"weixin://wxpay/bizpayurl?pr=qKcbbO7","status":200}
 *
 *
 *
 * @apiSuccess (Reponse 230) {Number} status  230
 * @apiSuccess (Reponse 230) {String} error 订单不存在
 *
 * @apiSuccess (Reponse 231) {Number} status  231
 * @apiSuccess (Reponse 231) {String} error 统一下单失败
 *
 * @apiSuccess (Reponse 410) {Number} status  410
 * @apiSuccess (Reponse 410) {String} error 错误参数
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
 * @apiSampleRequest http://121.41.116.104:8003/api/weixin/GetRedirectUrl
 */
//获取订单扫码url
func GetRedirectUrl(c *gin.Context) {
	order_info_id := c.Query("order_info_id")

	ret, code := ProductUnifiedOrder_pub("MWEB", order_info_id, "", "")
	if code == 230 {
		R.Error(c.Writer, 230, "订单不存在")
		return
	} else if code == 231 {
		R.Error(c.Writer, 231, "统一下单失败")
		return
	}

	R.Success(c.Writer, ret.Code_url)
}

/**
 * @api {get} /api/weixin/GetJsConfig 微信js支付参数
 * @apiName 微信js支付参数
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} order_info_id 订单id
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回扫码的链接
 * @apiSuccessExample Success-Response:
 * {"data":"weixin://wxpay/bizpayurl?pr=qKcbbO7","status":200}
 *
 *
 *
 * @apiSuccess (Reponse 230) {Number} status  230
 * @apiSuccess (Reponse 230) {String} error 订单不存在
 *
 * @apiSuccess (Reponse 231) {Number} status  231
 * @apiSuccess (Reponse 231) {String} error 统一下单失败
 *
 * @apiSuccess (Reponse 410) {Number} status  410
 * @apiSuccess (Reponse 410) {String} error 错误参数
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
 * @apiSampleRequest http://121.41.116.104:8003/api/weixin/GetJsConfig
 */
//获取订单扫码url
func GetJsConfig(c *gin.Context) {
	order_info_id := c.Query("order_info_id")
	token := c.Query("token")

	if token == "" || order_info_id == "" {
		R.Api404(c.Writer)
		return
	}

	r := redis.Get()
	defer r.Close()

	tokeninfo, _ := r.Get("baseweixinlogin:" + token)
	winfo := other.Json2Map(tokeninfo)

	resultmap, code := ProductUnifiedOrder_pub("JSAPI", order_info_id, winfo["data"].(map[string]interface{})["openid"].(string),
		strings.Split(c.Request.RemoteAddr, ":")[0])

	if code == 230 {
		R.Error(c.Writer, 230, "订单不存在")
		return
	} else if code == 231 {
		R.Error(c.Writer, 231, "统一下单失败")
		return
	}

	ret := map[string]interface{}{
		"appId":     w.Appid,
		"nonceStr":  w.Getnoncestr(),
		"timeStamp": fmt.Sprintf("%v", time.Now().Unix()),
		"package":   "prepay_id=" + resultmap.Prepay_id,
		"signType":  "MD5",
	}
	ret["paySign"] = w.PaySign(ret)
	ret["out_trade_no"] = order_info_id

	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data":   ret,
	})

}

// 微信统一下单
// 200 成功
// 230 订单不存在
// 231 统一下单失败
func ProductUnifiedOrder_pub(pay_type, order_info_id string, openid, spbill_create_ip string) (*wmsg.ResultXml, int) {
	order_info := product.ProductOrderInfoFind(order_info_id)
	if len(order_info) <= 0 || order_info["pay"] == "1" { //商品不存在
		return nil, 230
	}

	if pay_type == "JSAPI" {
		order_info_id = "WJ" + order_info_id
	} else if pay_type == "NATIVE" {
		order_info_id = "WN" + order_info_id
	} else if pay_type == "MWEB" {
		order_info_id = "WM" + order_info_id
	} else {
		order_info_id = "WO" + order_info_id
	}

	total_fee := other.ToFloat64(order_info["price"])
	data := map[string]interface{}{
		"body":         order_info["product_name"],
		"out_trade_no": order_info_id,
		"total_fee":    other.ToInt(total_fee * 100),
		"notify_url":   conf.GetString("config", "HOST_SERVER") + notify_url,
	}
	if pay_type == "JSAPI" {
		data["openid"] = openid
		data["spbill_create_ip"] = spbill_create_ip
	}
	resultmap := w.UnifiedOrder_pub(pay_type, data)
	fmt.Println(resultmap)
	if "FAIL" == resultmap.Return_code {
		return nil, 231
	}
	return resultmap, 200
}
