package product

import (
	"fmt"
	// "log"
	// "IginServer/conf"
	// "IginServer/lib/md5"
	"IginServer/lib/mysqldb"
	// "strings"
	"time"
)

func ProductFindAll() []map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("product").Where("del = 0").FindAll()

	return mysqldb.List(ret)
}

func ProductFind(productid string) map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("product").Where("del = 0 AND product_id = '" + productid + "'").FindOne()

	return ret[1]
}

// func ProductOrderFind(orderid string) map[string]string {
// 	db := mysqldb.GetConnect()

// 	ret := db.SetTable("product_order").Where("del = 0 AND order_id = '" + orderid + "'").FindOne()

// 	return ret[1]
// }

func ProductOrderInfoFind(infoid string) map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("product_order_info").Where("del = 0 AND order_info_id = '" + infoid + "'").FindOne()

	return ret[1]
}

// func CreateOrder(infoid, orderid, openid, pay_price, pay_type string) int {
// 	db := mysqldb.GetConnect()

// 	data := map[string]interface{}{
// 		"order_info_id": infoid,
// 		"order_id":      orderid,
// 		"weixin_openid": openid,
// 		"pay_price":     pay_price,
// 		"pay_type":      pay_type,
// 		"pay":           "0",
// 		"createtime":    time.Now().Unix(),
// 	}
// 	i, _ := db.IInsertOne("product_order", data)
// 	return int(i)
// }

func FinishOrder(order_no, pay_type, openid, transaction_id string, pay_price float64) int {
	db := mysqldb.GetConnect()

	data := map[string]interface{}{
		"pay":            "1",
		"transaction_id": transaction_id,
		"pay_price":      pay_price,
		"pay_type":       pay_type,
		"flow":           "2",
	}

	i, _ := db.SetTable("product_order_info").Where("del = 0 AND order_info_id = '" + order_no + "'").Update(data)
	return int(i)
}

func CreateOrderInfo(v map[string]interface{}) int {
	db := mysqldb.GetConnect()

	data := map[string]interface{}{
		"uid":                           v["uid"],
		"order_info_id":                 v["order_info_id"],
		"product_id":                    v["product_id"],
		"product_name":                  v["product_name"],
		"price":                         v["price"],
		"company":                       v["company"],
		"phone":                         v["phone"],
		"mail":                          v["mail"],
		"business_licence":              v["business_licence"],
		"tax_registration_certificate":  v["tax_registration_certificate"],
		"Organization_Code_Certificate": v["Organization_Code_Certificate"],
		"createtime":                    time.Now().Unix(),
	}
	i, _ := db.IInsertOne("product_order_info", data)
	// fmt.Println(err)
	return int(i)
}

//查询是否有未支付的订单产品
func FindUnPayOrder(uid string) map[string]string {
	db := mysqldb.GetConnect()

	where := "uid = '%s' AND del = 0 AND pay = 0"
	ret := db.SetTable("product_order_info").Where(fmt.Sprintf(where, uid)).FindOne()

	if len(ret) <= 0 {
		return nil
	} else {
		return ret[1]
	}
}
