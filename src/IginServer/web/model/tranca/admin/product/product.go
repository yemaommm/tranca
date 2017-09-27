package product

import (
	"fmt"
	// "log"
	"IginServer/conf"
	// "IginServer/lib/md5"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func ProductList(i, j int, search map[string]string) ([]map[string]string, string) {
	db := mysqldb.GetConnect()
	count := ""

	where := ""
	if search["name"] != "" {
		where += fmt.Sprintf(" AND name LIKE '%s' ", "%"+search["name"]+"%")
	}

	ret, _ := db.IExecute(fmt.Sprintf(conf.XmlGet("tranca.product.findAllList"), where), i, j)
	db.SetTable("product a").Fileds("COUNT(1) AS n")
	if where != "" {
		where = where[4:]
		db.Where(where)
	}
	count = db.FindAll()[1]["n"]

	return ret, count
}

func DelProduct(product_id, aid string) int {
	db := mysqldb.GetConnect()
	i, _ := db.SetTable("product").Where("aid = '" + aid + "' and product_id = '" + product_id + "'").Update(map[string]interface{}{"del": 1})

	return i
}

func AddProduct(param map[string]interface{}) int {
	db := mysqldb.GetConnect()

	i, err := db.IInsertOne("product", param)
	fmt.Println(err)
	return int(i)
}

func UpdateProduct(aid, product_id string, param map[string]interface{}) int {
	db := mysqldb.GetConnect()
	//+ "' and aid = '" + aid
	i, _ := db.SetTable("product").Where("product_id = '" + product_id + "' AND del = 0").Update(param)
	return i
}

func ProductInfo(aid, product_id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("product").Where("product_id = '" + product_id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret
}
