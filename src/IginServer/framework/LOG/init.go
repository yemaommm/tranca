package LOG

import (
	// "log"
	// "server/framework/API/I"
	"IginServer/lib/Imartini"
	// "net/http"
)

// var base = "M"
// var apiRoute = [][]interface{}{
// 	{"M/log", "POST", senddoc["url"], &I.Atype{senddoc, send}}, //有新订单时进行广播
// }
// var route = [][]interface{}{
// 	{"", "GET", "/log/test", test},               //有新订单时进行广播
// 	{"M/log", "POST", "/log/sendjson", sendjson}, //有新订单时进行广播
// }

func init() {
	// Imartini.M.AddIroute(base, route)
	// Imartini.M.AddIroute(base, apiRoute, "api")
	// Imartini.M.AddRoute("GET", "/js/", http.FileServer(http.Dir("/var/www/js/")))
	// log.Println(Imartini.API_URL)
	router := Imartini.M.Group("M")
	router.GET("/log/test", test)
	router.POST("/log/sendjson", sendjson)
	router.POST("/log/send", send)
}
