package weixin

import (
	// "github.com/martini-contrib/render"
	// "log"
	// "github.com/go-martini/martini"
	"IginServer/conf"
	"IginServer/framework/API/R"
	"IginServer/framework/weixin/msg"
	"IginServer/lib/Imartini"
	"IginServer/lib/other"
	"IginServer/lib/redis"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	Login_wx_url     = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
	Login_basewx_url = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"

	Access_token_url = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	Jsapi_ticket_url = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

	User_access_token_url   = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	User_reaccess_token_url = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	User_userinfo_url       = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	User_auth               = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	User_sub                = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

var (
	Appid             = "wx45725ee97184f63a"               //"wx6c8cc65a047c52e7"
	AppSecret         = "f6f5b4aa892412e640101d8be5f4102a" //"ae9f9a5988876cb9ece224dae0de6fbe"
	Mch_id            = "1303712101"
	APP_KEY           = "shanghaichuankewangluokejidoubix"
	Access_token      = ""
	Jsapi_ticket      = ""
	Access_token_time = 0
	Jsapi_ticket_time = 0
)

/**
 * @api {get} /weixin/sub 判断用户是否关注
 * @apiName 判断用户是否关注
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} openid openid
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/sub
 */
/*  获取用户信息 */
func Sub(c *gin.Context) {
	// res.Header().Add("Access-Control-Allow-Origin", "*")
	if c.Request.FormValue("openid") == "" {
		R.Api404(c.Writer)
		return
	}
	getaccess_token()
	ret := other.Httpget(fmt.Sprintf(User_sub, Access_token, c.Request.FormValue("openid")))
	// log.Printf("%s", ret)
	var v map[string]interface{}
	json.Unmarshal(ret, &v)
	c.JSON(200, map[string]interface{}{
		"status": 200,
		"data":   v,
	})
}

/**
 * @api {get} /weixin/basecode 获取基础用户信息
 * @apiName 获取基础用户信息
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} code code
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
 *
 * @apiSuccess (Reponse 300) {Number} status  300
 * @apiSuccess (Reponse 300) {String} error access_token获取失败
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/basecode
 */
/*  获取用户信息 */
func Basecode(c *gin.Context) {
	// res.Header().Add("Access-Control-Allow-Origin", "*")
	if c.Request.FormValue("code") == "" {
		R.Api404(c.Writer)
		return
	}
	var v map[string]interface{}
	code := c.Request.FormValue("code")
	// ntime := int(time.Now().Unix())
	stmp := other.Httpget(fmt.Sprintf(User_access_token_url, Appid, AppSecret, code))
	json.Unmarshal(stmp, &v)

	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data":   v,
	})

}

/**
 * @api {get} /weixin/ApiBaseWeixinLogin 作为借口的微信登录
 * @apiName 作为借口的微信登录
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} token token
 * @apiParam {String} re_url 微信登录完之后跳回的url，选填
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 * 跳转完这个链接之后，需要查询微信返回的数据，用token调用另一个借口查询
 *
 *
 *
 * @apiSuccess (Reponse 300) {Number} status  300
 * @apiSuccess (Reponse 300) {String} error access_token获取失败
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/ApiBaseWeixinLogin
 */
func ApiBaseWeixinLogin(c *gin.Context) {
	re_url := c.Request.FormValue("re_url")
	token := c.Request.FormValue("token")
	r := redis.Get()
	defer r.Close()

	tokeninfo, _ := r.Get("baseweixinlogin:" + token)

	if tokeninfo == "" && strings.Index(c.Request.Header.Get("User-Agent"), "MicroMessenger") == -1 {
		// sess["last_url"] = c.Request.Header.Get("Referer")
		c.String(200, "请用微信访问")
		return
	} else if c.Request.FormValue("code") == "" && tokeninfo == "" {
		// sess["last_url"] = c.Request.Header.Get("Referer")

		now_url := "http://" + c.Request.Host + c.Request.URL.RequestURI()
		wxurl := fmt.Sprintf(Login_basewx_url, Appid, url.QueryEscape(now_url))

		// fmt.Printf("%s\n", other.HttpCopyGet(c.Request, wxurl))
		http.Redirect(c.Writer, c.Request, wxurl, http.StatusFound)
		return
	} else if c.Request.FormValue("code") != "" && tokeninfo == "" {
		// weixinconfig := weixin.Weixininfo("0")

		// userinfo := GetUserInfo(weixinconfig["appid"], weixinconfig["appsecret"], c.Request.FormValue("code"))
		stmp := other.Httpget(conf.GetString("config", "HOST_API") + "weixin/basecode?code=" + c.Request.FormValue("code"))

		tokeninfo = string(stmp)
		r.Set("baseweixinlogin:"+token, stmp)
		r.Expire("baseweixinlogin:"+token, 7200)
	}

	if re_url != "" {
		if len(strings.Split(re_url, "?")) < 2 {
			re_url += "?"
		}
		if string(re_url[len(re_url)-1]) != "&" && string(re_url[len(re_url)-1]) != "?" {
			re_url += "&"
		}
		re_url += "base=ApiBaseWeixinLogin"
		http.Redirect(c.Writer, c.Request, re_url, http.StatusFound)
	}
}

/**
 * @api {get} /weixin/GetWeixinLoginInfo 获取微信登录信息
 * @apiName 获取微信登录信息
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} token token
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
 * @apiSuccess (Reponse 300) {Number} status  300
 * @apiSuccess (Reponse 300) {String} error access_token获取失败
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/GetWeixinLoginInfo
 */
func GetWeixinLoginInfo(c *gin.Context) {
	token := c.Request.FormValue("token")
	r := redis.Get()
	defer r.Close()

	tokeninfo, _ := r.Get("baseweixinlogin:" + token)

	R.Success(c.Writer, other.Json2Map(tokeninfo))
}

/**
 * @api {get} /weixin/code 获取用户信息
 * @apiName 获取用户信息
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} code code
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
 *
 * @apiSuccess (Reponse 300) {Number} status  300
 * @apiSuccess (Reponse 300) {String} error access_token获取失败
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/code
 */
/*  获取用户信息 */
func Code(c *gin.Context) {
	// res.Header().Add("Access-Control-Allow-Origin", "*")
	if c.Request.FormValue("code") == "" {
		R.Api404(c.Writer)
		return
	}
	var v map[string]interface{}
	var uinfo map[string]interface{}
	code := c.Request.FormValue("code")
	// ntime := int(time.Now().Unix())
	stmp := other.Httpget(fmt.Sprintf(User_access_token_url, Appid, AppSecret, code))
	json.Unmarshal(stmp, &v)
	Imartini.MyLog.Printf("%v", v)
	if v["errcode"] != nil {
		R.Write(c.Writer, map[string]interface{}{
			"status": 300,
			"error":  "access_token获取失败",
		})
		return
	}
	//获取用户信息
	istmp := other.Httpget(fmt.Sprintf(User_userinfo_url, v["access_token"].(string), v["openid"].(string)))
	json.Unmarshal(istmp, &uinfo)
	if uinfo["errcode"] != nil {
		Imartini.MyLog.Printf("%s", istmp)
		R.Write(c.Writer, map[string]interface{}{
			"status": 200,
			"data":   v,
		})
		return
	}
	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data":   uinfo,
	})

}

/**
 * @api {get} /weixin/info 获取wx.config参数
 * @apiName 获取wx.config参数
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} url url
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/info
 */
/*  wx.config提供的参数 */
func Info(c *gin.Context) {
	// res.Header().Add("Access-Control-Allow-Origin", "*")
	if c.Request.FormValue("url") == "" {
		R.Api404(c.Writer)
		return
	}
	// fmt.Println(c.Request.FormValue("url"))
	getaccess_token()
	signature := map[string]interface{}{
		"jsapi_ticket": Jsapi_ticket,
		"noncestr":     Getnoncestr(),
		"timestamp":    int(time.Now().Unix()),
		"url":          c.Request.FormValue("url"),
	}

	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"appId":     Appid,
			"timestamp": signature["timestamp"],
			"noncestr":  signature["noncestr"],
			"signature": Sign(signature),
		},
	})
}

/**
 * @api {get} /weixin/NativeLink 获取静态扫码支付URL
 * @apiName 获取静态扫码支付URL
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} productid productid
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/NativeLink
 */
func NativeLink(c *gin.Context) {
	productid := c.Query("productid")

	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data": map[string]interface{}{
			"url": NativeURL(productid),
		},
	})
}

/**
 * @api {post get} /weixin/jspay 获取JS支付参数
 * @apiName 获取JS支付参数
 * @apiGroup 微信
 *
 * @apiVersion 1.0.0
 *
 * @apiParam {String} product_body product_body
 * @apiParam {String} out_trade_no out_trade_no
 * @apiParam {String} total_fee total_fee
 * @apiParam {String} spbill_create_ip spbill_create_ip
 * @apiParam {String} notify_url notify_url
 * @apiParam {String} openid openid
 *
 * @apiSuccess (Reponse 200) {Number} status 200
 * @apiSuccess (Reponse 200) {String} data  返回信息
 * @apiSuccessExample Success-Response:
 *
 *
 *
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
 * @apiSampleRequest http://121.41.116.104:8003/weixin/jspay
 */
func JsPay(c *gin.Context) {
	// reqHandler.init();
	// reqHandler.init(validate.appid, validate.appSecret, validate.APP_KEY, validate.PARTNER_KEY);
	// //执行统一下单接口 获得预支付id
	// reqHandler.setParameter("appid",validate.appid);
	// reqHandler.setParameter("mch_id", validate.MCHID);                //商户号
	// reqHandler.setParameter("nonce_str", noncestr);            //随机字符串
	// reqHandler.setParameter("body", product_name);                        //商品描述(必填.如果不填.也会提示系统升级.正在维护我艹.)
	// reqHandler.setParameter("out_trade_no", out_trade_no);        //商家订单号
	// reqHandler.setParameter("total_fee", order_price);                    //商品金额,以分为单位
	// reqHandler.setParameter("spbill_create_ip", http.getRealIP(req));   //用户的公网ip  IpAddressUtil.getIpAddr(request)
	// //下面的notify_url是用户支付成功后为微信调用的action  异步回调.
	// reqHandler.setParameter("notify_url", validate.NOTIFY_URL+"/"+id);
	// reqHandler.setParameter("trade_type", "JSAPI");
	// //------------需要进行用户授权获取用户openid-------------
	// reqHandler.setParameter("openid", openid);   //这个必填.
	product_name := c.Request.FormValue("product_body")
	out_trade_no := c.Request.FormValue("out_trade_no")
	total_fee := c.Request.FormValue("total_fee")
	spbill_create_ip := c.Request.FormValue("spbill_create_ip")
	notify_url := c.Request.FormValue("notify_url")
	openid := c.Request.FormValue("openid")

	// data := map[string]interface{}{
	// 	"appid":            Appid,
	// 	"mch_id":           Mch_id,
	// 	"nonce_str":        Getnoncestr(),
	// 	"body":             product_name,
	// 	"out_trade_no":     out_trade_no,
	// 	"total_fee":        total_fee,
	// 	"spbill_create_ip": spbill_create_ip,
	// 	"notify_url":       notify_url,
	// 	"trade_type":       "JSAPI",
	// 	"openid":           openid,
	// }
	// data["sign"] = PaySign(data)

	// // fmt.Println(msg.MapToXml(data))
	// // fmt.Println(string(other.HttpPostBody("https://api.mch.weixin.qq.com/pay/unifiedorder", []byte(msg.MapToXml(data)))))
	// result := string(other.HttpPostBody("https://api.mch.weixin.qq.com/pay/unifiedorder", []byte(msg.MapToXml(data))))
	// resultmap := new(msg.ResultXml)
	// xml.Unmarshal([]byte(result), &resultmap)

	data := map[string]interface{}{
		"body":             product_name,
		"out_trade_no":     out_trade_no,
		"total_fee":        total_fee,
		"spbill_create_ip": spbill_create_ip,
		"notify_url":       notify_url,
		"openid":           openid,
	}
	resultmap := UnifiedOrder_pub("JSAPI", data)

	if "FAIL" == resultmap.Return_code {
		R.Error(c.Writer, 410, resultmap.Return_msg)
		return
	}

	ret := map[string]interface{}{
		"appId":     Appid,
		"nonceStr":  Getnoncestr(),
		"timeStamp": fmt.Sprintf("%v", time.Now().Unix()),
		"package":   "prepay_id=" + resultmap.Prepay_id,
		"signType":  "MD5",
	}
	ret["paySign"] = PaySign(ret)
	ret["out_trade_no"] = out_trade_no

	R.Write(c.Writer, map[string]interface{}{
		"status": 200,
		"data":   ret,
	})
}

//静态扫码URL
func NativeURL(productid string) string {
	data := map[string]interface{}{
		"appid":      Appid,
		"mch_id":     Mch_id,
		"time_stamp": fmt.Sprintf("%v", time.Now().Unix()),
		"nonce_str":  Getnoncestr(),
		"product_id": productid,
	}
	data["sign"] = PaySign(data)
	bizstring := FormatBizQueryParaMap(data)

	return "weixin://wxpay/bizpayurl?" + bizstring
}

//同一下单
// "body":             product_name,
// "out_trade_no":     out_trade_no,
// "total_fee":        total_fee,
// "spbill_create_ip": spbill_create_ip,
// "notify_url":       notify_url,
// "openid":           openid,  //trade_type为jsapi必须要opendid
func UnifiedOrder_pub(trade_type string, data map[string]interface{}) *msg.ResultXml {
	data["appid"] = Appid
	data["mch_id"] = Mch_id
	data["nonce_str"] = Getnoncestr()
	data["trade_type"] = trade_type
	data["sign"] = PaySign(data)

	// fmt.Println(msg.MapToXml(data))
	// fmt.Println(string(other.HttpPostBody("https://api.mch.weixin.qq.com/pay/unifiedorder", []byte(msg.MapToXml(data)))))
	result := string(other.HttpPostBody("https://api.mch.weixin.qq.com/pay/unifiedorder", []byte(msg.MapToXml(data))))
	fmt.Println(result)
	resultmap := new(msg.ResultXml)
	xml.Unmarshal([]byte(result), &resultmap)

	return resultmap
}

func FormatBizQueryParaMap(signature map[string]interface{}) string {
	signature_key := make([]string, 0)
	signature_str := ""
	for i, _ := range signature {
		signature_key = append(signature_key, i)
	}
	sort.Strings(signature_key)
	for _, i := range signature_key {
		signature_str += i + "=" + fmt.Sprintf("%v", signature[i]) + "&"
	}
	signature_str = signature_str[0 : len(signature_str)-1]
	return signature_str
}

func PaySign(signature map[string]interface{}) string {
	signature_str := FormatBizQueryParaMap(signature)
	signature_str += "&key=" + APP_KEY
	fmt.Println(signature_str)
	str := fmt.Sprintf("%x", md5.Sum([]byte(signature_str)))
	str = strings.ToUpper(str)
	fmt.Println(str)

	return str
}

func Sign(signature map[string]interface{}) string {
	signature_str := FormatBizQueryParaMap(signature)
	fmt.Println(signature_str)
	str := fmt.Sprintf("%x", sha1.Sum([]byte(signature_str)))
	fmt.Println(str)

	return str
}

func getaccess_token() {
	if int(time.Now().Unix())-Access_token_time >= 7100 {
		Access_token_time = int(time.Now().Unix())
		access := other.Httpget(fmt.Sprintf(Access_token_url, Appid, AppSecret))
		fmt.Printf("%s\n", access)
		var v map[string]interface{}
		json.Unmarshal(access, &v)
		Access_token = fmt.Sprintf("%v", v["access_token"])
	}
	if int(time.Now().Unix())-Jsapi_ticket_time >= 7100 {
		Jsapi_ticket_time = int(time.Now().Unix())
		stmp := other.Httpget(fmt.Sprintf(Jsapi_ticket_url, Access_token))
		fmt.Printf("%s\n", stmp)
		var v map[string]interface{}
		json.Unmarshal(stmp, &v)
		Jsapi_ticket = fmt.Sprintf("%v", v["ticket"])
	}
}

func Getnoncestr() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v", int(time.Now().Unix())))))
}
