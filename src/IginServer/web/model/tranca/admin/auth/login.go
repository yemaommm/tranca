package auth

import (
	"fmt"
	// "log"
	"IginServer/lib/mysqldb"
	// "encoding/json"
)

func Login(username, password string) (int, map[string]string) {
	sql := `SELECT a.id, a.username, a.passwd
        , a.token, a.tel, a.group, b.qrcode 
        FROM admin_auth a 
        LEFT JOIN admin_group b 
        ON a.group = b.id 
        WHERE a.username = '%v'
        AND a.passwd = '%v'`
	db := mysqldb.GetConnect()

	ret := db.Query(fmt.Sprintf(sql, username, password))

	if len(ret.(map[int]map[string]string)) <= 0 {
		return 0, map[string]string{}
	} else {
		// js, _ := json.Marshal(ret.(map[int]map[string]string)[1])
		return 1, ret.(map[int]map[string]string)[1]
	}
}
