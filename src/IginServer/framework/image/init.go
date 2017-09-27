package image

import (
	// "log"
	"IginServer/lib/Imartini"
	// "net/http"
)

// var base = "M"
// var route = [][]interface{}{
// 	{"", "POST", "/file/upload", uploadfile},        //文件上传
// 	{"", "POST", "/image/upload", postimg},          //图片上传
// 	{"", "GET", "/upload/(?P<path>[^#?]*)", catimg}, //图片上传
// }

func init() {
	// Imartini.M.AddIroute(base, route)
	// Imartini.M.AddRoute("GET", "/js/", http.FileServer(http.Dir("/var/www/js/")))
	// log.Println(Imartini.API_URL)
	router := Imartini.M.Group("/M")
	router.POST("/file/upload", uploadfile) //文件上传
	router.POST("/image/upload", postimg)   //图片上传
	// router.GET("/upload/(?P<path>[^#?]*)", catimg) //查看图片
	router.GET("/upload/*path", catimg) //查看图片
}
