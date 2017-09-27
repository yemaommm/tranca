package auth

import (
	// "log"
	"IginServer/lib/mysqldb"
)

func Register(param map[string]interface{}) error {
	db := mysqldb.GetConnect()
	i, err := db.SetTable("admin_group").Insert(map[string]interface{}{"qrcode": 0})
	if err != nil {
		return err
	}
	param["group"] = i
	_, err = db.SetTable("admin_auth").Insert(param)
	if err != nil {
		return err
	}
	return nil
}

func Isexist(username string) int {
	db := mysqldb.GetConnect()
	ret := db.SetTable("admin_auth").Where("username='" + username + "'").FindOne()

	if len(ret) > 0 {
		return 1
	}
	return 0
}
