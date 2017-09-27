package weixin

import (
	"fmt"
	// "log"
	// "server/conf"
	w "IginServer/framework/weixin"
	"IginServer/lib/mysqldb"
	// "strconv"
	"time"
)

func Weixininfo(uid string) map[string]string {
	db := mysqldb.GetConnect()

	// where := fmt.Sprintf("uid = %s", uid).Where(where)

	db.SetTable("weixin_setting").Limit(1)

	data := db.FindOne()
	return data[1]
}

func UpdateWeixininfo(uid string, data map[string]interface{}) {
	db := mysqldb.GetConnect()

	// where := fmt.Sprintf("uid = %s", uid)//.Where(where)
	data["uid"] = uid
	data["updatetime"] = int(time.Now().Unix())

	if i, _ := db.SetTable("weixin_setting").Update(data); i <= 0 {
		data["uid"] = uid
		db.SetTable("weixin_setting").Insert(data)
	}

}

func SetWeixinConfig() {
	info := Weixininfo("0")

	w.Appid = info["appid"]
	w.AppSecret = info["appsecret"]
	w.Mch_id = info["mchid"]
	w.APP_KEY = info["key"]

	fmt.Println(w.Appid, w.AppSecret, w.Mch_id, w.APP_KEY)
}
