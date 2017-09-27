package wap_user

import (
	// "fmt"
	// "log"
	// "IginServer/conf"
	"IginServer/lib/md5"
	"IginServer/lib/mysqldb"
	// "strings"
	"time"
)

func Login(username, password string) map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("wap_user").Fileds("username", "uid").
		Where("del = 0 AND username = '" + username + "' AND password = '" + md5.Md5(password) + "'").FindOne()

	return ret[1]
}

func Register(username, password string) int {
	db := mysqldb.GetConnect()

	data := map[string]interface{}{
		"username":   username,
		"password":   md5.Md5(password),
		"uid":        md5.GetUUID().Hex(),
		"createtime": time.Now().Unix(),
	}
	i, _ := db.IInsertOne("wap_user", data)

	return int(i)
}

func FindUserForUsername(username string) map[string]string {
	db := mysqldb.GetConnect()

	ret := db.SetTable("wap_user").Where("del = 0 AND username = '" + username + "'").FindOne()

	if len(ret) <= 0 {
		return nil
	}
	return ret[1]
}
