package mysqldb

import (
	// "database/sql"
	// "fmt"
	"IginServer/conf"
	"IginServer/lib/mysqldb/obj"
	"strconv"
)

var xdb *obj.MyDB

func GetConnect() *obj.DB {
	// err := xdb.conn.Ping()
	// if err != nil {
	// 	New()
	// }
	return obj.CreateMyDB(xdb)
}

func GetNew(db *obj.MyDB) *obj.DB {
	return obj.CreateMyDB(db)
}

func New() {
	DB_DRIVE := conf.GET["config"]["DB_DRIVE"]
	DB_MAX_OPEN, _ := strconv.Atoi(conf.GET["config"]["DB_MAX_OPEN"])
	DB_MAX_IDLE, _ := strconv.Atoi(conf.GET["config"]["DB_MAX_IDLE"])
	DB_NAME := conf.GET["config"]["DB_NAME"]
	DB_HOST := conf.GET["config"]["DB_HOST"]
	DB_USER := conf.GET["config"]["DB_USER"]
	DB_PASSWD := conf.GET["config"]["DB_PASSWD"]
	DB_PORT := conf.GET["config"]["DB_PORT"]
	DB_CONN := DB_USER + ":" + DB_PASSWD + "@(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8"

	xdb = obj.Create(DB_DRIVE, DB_CONN, DB_MAX_IDLE, DB_MAX_OPEN)
}

func List(i interface{}) []map[string]string {
	stmp := i.(map[int]map[string]string)
	ret := make([]map[string]string, 0)

	for i := 1; i <= len(stmp); i++ {
		ret = append(ret, stmp[i])
	}
	return ret
}

func ToMap(list []map[string]string) map[int]map[string]string {
	ret := make(map[int]map[string]string)

	for i, j := range list {
		ret[i+1] = j
	}
	return ret
}

func DBdefer(db *obj.Transaction) {
	if err := recover(); err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}

func init() {
	New()
}
