package weixin

import "IginServer/lib/Imartini"

// // import "log"
// // import "net/http"

// var base = "weixin"
// var route = [][]interface{}{
// 	{"", "GET", "/sub", sub},   //判断用户是否关注
// 	{"", "GET", "/code", code}, //获取用户信息
// 	{"", "GET", "/info", info}, //获取wx.config参数
// 	{"", "*", "/pay", pay},     //获取JS支付参数
// 	{"", "GET", "/validate", validate},
// 	{"", "POST", "/validate", pvalidate},
// }

func init() {
	// Imartini.M.AddIroute(base, route)
	// Imartini.M.AddRoute("GET", "/js/", http.FileServer(http.Dir("/var/www/js/")))
	// log.Println(Imartini.M.Iroute)
	router := Imartini.M.Group("/weixin")
	router.GET("/sub", Sub)           //判断用户是否关注
	router.GET("/code", Code)         //获取用户信息
	router.GET("/basecode", Basecode) //获取用户信息
	router.GET("/info", Info)         //获取wx.config参数
	router.GET("/jspay", JsPay)       //获取JS支付参数
	router.POST("/jspay", JsPay)      //获取JS支付参数
	router.GET("/validate", validate)
	router.POST("/validate", pvalidate)
	router.GET("/NativeLink", NativeLink) //获取扫码支付URL

	router.GET("/ApiBaseWeixinLogin", ApiBaseWeixinLogin) //作为借口的微信登录
	router.GET("/GetWeixinLoginInfo", GetWeixinLoginInfo) //获取微信登录信息
}
