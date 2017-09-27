package Imartini

// import (
// 	"github.com/go-martini/martini"
// 	"net/http"
// 	"server/lib/mysqldb"
// )

// func DB() martini.Handler {
// 	return func(res http.ResponseWriter, req *http.Request, c martini.Context, mlog *Mlog) {
// 		db := mysqldb.GetConnect()
// 		tarn, _ := db.Bagin()
// 		c.Map(tarn)
// 		c.Next()
// 		tarn.Commit()
// 	}
// }
