package order

import (
	"fmt"
	// "log"
	"IginServer/conf"
	// "IginServer/lib/md5"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func OrderList(i, j int, search map[string]string) ([]map[string]string, string) {
	db := mysqldb.GetConnect()
	count := ""

	where := ""
	for i, j := range search {
		if j != "" && j != "-1" {
			where += " AND " + i + "='" + j + "'"
		}
	}

	ret, _ := db.IExecute(fmt.Sprintf(conf.XmlGet("tranca.order.findAllList"), where), i, j)
	db.SetTable("product_order_info a").LeftJoin("wap_user b", "a.uid = b.uid").Fileds("COUNT(1) AS n")
	if where != "" {
		where = where[4:]
		db.Where(where)
	}
	count = db.FindAll()[1]["n"]

	return ret, count
}

func DelOrder(order_info_id, aid string) int {
	db := mysqldb.GetConnect()
	i, _ := db.SetTable("product_order_info").Where(" order_info_id = '" + order_info_id + "'").
		Update(map[string]interface{}{"del": aid, "aid": aid})

	return i
}

func AddOrder(param map[string]interface{}) int {
	db := mysqldb.GetConnect()

	i, err := db.IInsertOne("product_order_info", param)
	fmt.Println(err)
	return int(i)
}

func UpdateOrder(aid, order_info_id string, param map[string]interface{}) int {
	db := mysqldb.GetConnect()
	param["aid"] = aid
	i, _ := db.SetTable("product_order_info").Where("order_info_id = '" + order_info_id + "' AND del = 0").Update(param)
	return i
}

func OrderInfo(aid, order_info_id string) map[int]map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("product_order_info").Where(" order_info_id = '" + order_info_id + "'").FindOne()

	// top := []map[string]string{}
	// for _, i := range ret {
	//  if i["level"] == "1" {
	//      top = append(top, i)
	//  }
	// }
	// fmt.Println(top)

	return ret
}
