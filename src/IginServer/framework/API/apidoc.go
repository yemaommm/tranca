package API

// import (
// 	"github.com/martini-contrib/render"
// 	"log"
// 	"net/http"
// 	"server/conf"
// 	"server/lib/Imartini"
// 	"strconv"
// )

// var CATURL, _ = strconv.ParseBool(conf.GET["config"]["CATURL"])

// func Create() {
// 	for i, j := range Imartini.API_URL {
// 		base := i
// 		body := j
// 		Imartini.M.Get("/api/doc/"+base, func(r render.Render, res http.ResponseWriter, req *http.Request, log *Imartini.Mlog) {
// 			count := 1
// 			for x, y := range body {
// 				if req.FormValue("type") == x || len(body) == count {
// 					TYPE := x
// 					DATA := y
// 					stmp := Imartini.API_URL[base][TYPE]
// 					for i, _ := range stmp {
// 						stmp[i].(map[string]interface{})["host"] = req.Host
// 					}
// 					r.HTML("framework/apidoc/doc.html", map[string]interface{}{
// 						"data": DATA,
// 						"type": body,
// 						// "ids":  req.FormValue("ids"),
// 					})
// 					return
// 				}
// 				count++
// 			}
// 		})
// 		if CATURL {
// 			log.Printf("|%-40v|%-10v|%-40v|", "", "GET", "api/doc/"+base)
// 		}
// 	}
// 	Imartini.M.Get("/api/doc", func(r render.Render, res http.ResponseWriter, req *http.Request, log *Imartini.Mlog) {
// 		r.HTML("framework/apidoc/menu.html", map[string]interface{}{
// 			"data": Imartini.API_URL,
// 		})
// 	})
// }

// func init() {
// 	Create()
// }
