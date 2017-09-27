package weixin

import (
	// "IginServer/lib/Imartini"
	// "IginServer/lib/mygin"
	// "IginServer/lib/mysqldb"
	// "IginServer/conf"
	"IginServer/web/model/tranca/product"
	// "encoding/json"
	w "IginServer/framework/weixin"
	// "IginServer/lib/md5"
	"IginServer/lib/other"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	// "net/http"
	"time"
)

const (
	notify_url = "/api/weixin/NativePayNotify"
)

//扫码回调
func NativeNotify(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	fmt.Printf("%s\n", body)
	payresult := other.XML2Map(body, "xml")
	fmt.Println(payresult)

	sign := payresult["sign"].(string)
	delete(payresult, "sign")
	if w.PaySign(payresult) != sign { //签名不正确
		fmt.Println("微信支付：", "签名不正确")
		c.String(200, "<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[]]></return_msg></xml>")
		return
	}

	resultmap, code := ProductUnifiedOrder_pub("NATIVE", payresult["product_id"].(string), "", "")
	if code == 230 {
		c.String(200, "<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[]]></return_msg></xml>")
		return
	} else if code == 231 {
		c.String(200, "<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[]]></return_msg></xml>")
		return
	}

	parameters := map[string]interface{}{
		"return_code": "SUCCESS",
		"appid":       w.Appid,
		"mch_id":      w.Mch_id,
		"nonce_str":   w.Getnoncestr(),
		"time_stamp":  time.Now().Unix(),
		"prepay_id":   resultmap.Prepay_id,
		"result_code": "SUCCESS",
	}
	parameters["sign"] = w.PaySign(parameters)
	c.String(200, other.Map2XML(parameters))

	// if "FAIL" == resultmap.Return_code {
	//     R.Error(c.Writer, 410, resultmap.Return_msg)
	//     return
	// }
}

//扫码支付后回调
func NativePayNotify(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	fmt.Printf("%s\n", body)
	payresult := other.XML2Map(body, "xml")
	fmt.Println(payresult)

	if payresult["return_code"].(string) == "SUCCESS" { // 支付成功
		sign := payresult["sign"].(string)
		delete(payresult, "sign")
		if w.PaySign(payresult) != sign { //签名不正确
			fmt.Println("微信支付：", "签名不正确")
			return
		}

		total_fee := other.ToFloat64(payresult["total_fee"].(string))
		total_fee = total_fee / 100
		// fmt.Println("Total_fee", total_fee)
		// fmt.Println("Out_trade_no", payresult["Out_trade_no)
		product.FinishOrder(payresult["out_trade_no"].(string)[2:], "weixin", payresult["openid"].(string), payresult["transaction_id"].(string), total_fee)
	} else {
		c.String(200, "<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[]]></return_msg></xml>")
		return
	}
	c.String(200, "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[]]></return_msg></xml>")
}

func init() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(md5.GetUUID().Hex())
	// 	time.Sleep(1 * time.Second)
	// }

	// parameters := map[string]interface{}{
	// 	"return_code": "SUCCESS",
	// 	"appid":       123124,
	// 	"mch_id":      "123123",
	// 	"result_code": "SUCCESS",
	// }
	// fmt.Println(other.Map2XML(parameters))
}
