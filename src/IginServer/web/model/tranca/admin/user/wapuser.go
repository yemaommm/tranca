package user

import (
	"fmt"
	// "log"
	"IginServer/conf"
	"IginServer/lib/md5"
	"IginServer/lib/mysqldb"
	// "strconv"
)

func WapUserList(i, j int, search map[string]string) ([]map[string]string, string) {
	db := mysqldb.GetConnect()
	count := ""

	where := ""
	if search["username"] != "" {
		where += fmt.Sprintf(" AND username = '%s' ", search["username"])
	}

	ret, _ := db.IExecute(fmt.Sprintf(conf.XmlGet("tranca.user.finduserlist"), where), i, j)
	db.SetTable("wap_user a").Fileds("COUNT(1) AS n")
	if where != "" {
		where = where[4:]
		db.Where(where)
	}
	count = db.FindAll()[1]["n"]

	return ret, count
}

func UpdateWapUserPassword(uid, password string) int {
	db := mysqldb.GetConnect()

	i, _ := db.SetTable("wap_user").Where(fmt.Sprintf("uid = '%s'", uid)).Update(map[string]interface{}{"password": md5.Md5(password)})

	return i
}
