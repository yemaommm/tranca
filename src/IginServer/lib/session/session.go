package session

import (
	"fmt"
	// "github.com/go-martini/martini"
	"IginServer/conf"
	"IginServer/lib/cookie"
	"IginServer/lib/md5"
	"IginServer/lib/other"
	"IginServer/lib/redis"
	// "github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	SESSIONID       = conf.GET["config"]["SESSIONID"]
	SESSION_TIME, _ = strconv.Atoi(conf.GET["config"]["SESSION_TIME"])
	SESSION_KEY     = conf.GET["config"]["SESSION_KEY"]
	SESSION_TYPE    = conf.GET["config"]["SESSION_TYPE"]
)

type SESS map[string]interface{}

// func SESSION() martini.Handler {
// 	return func(res http.ResponseWriter, req *http.Request, r *redis.Redis, c martini.Context) {
// 		co, err := req.Cookie(SESSIONID)
// 		isid := ""
// 		if err != nil {
// 			t := time.Now().Unix()
// 			value := fmt.Sprintf("SESSION:%d|%s", t, req.RemoteAddr)
// 			cookie.Set(res, SESSIONID, value)
// 			isid = value
// 		} else {
// 			isid = co.Value
// 		}
// 		// fmt.Println("isid")
// 		// fmt.Println(isid)

// 		if SESSION_TYPE == "redis" {
// 			var ret map[string]string

// 			ret, _ = r.HGetAll(isid)

// 			//把查询出来的session复制一份出来，用于之后比对是否有变化
// 			old := make(map[string]string)
// 			for i, j := range ret {
// 				old[i] = j
// 			}
// 			//把session映射到请求
// 			c.Map(SESS(ret))
// 			//把redis映射到请求

// 			c.Next()

// 			//最后处理
// 			var value []interface{}
// 			var void []interface{}
// 			value = append(value, isid)
// 			void = append(void, isid)
// 			for i, j := range ret {
// 				//比对session中是否有变化
// 				if old[i] != j {
// 					value = append(value, i, j)
// 				}
// 				delete(old, i)
// 			}
// 			// 删除以删除的session
// 			for i, _ := range old {
// 				void = append(void, i)
// 			}

// 			if len(void) > 1 {
// 				r.HDel(void...)
// 			}
// 			if len(value) > 1 {
// 				r.HMSet(value...)
// 				r.Expire(isid, SESSION_TIME)
// 			}

// 		} else if SESSION_TYPE == "auto" {
// 			if autoSession[isid] == nil {
// 				autoSession[isid] = make(SESS)
// 				autoSessiontime[isid] = time.Now()
// 			}
// 			c.Map(autoSession[isid])
// 			c.Next()

// 			for i, j := range autoSessiontime {
// 				if float64(SESSION_TIME) < time.Since(j).Seconds() {
// 					delete(autoSessiontime, i)
// 				}
// 			}
// 		}
// 	}
// }

func GetSessionID(res http.ResponseWriter, req *http.Request) string {
	co, err := req.Cookie(SESSIONID)
	id := ""
	if err != nil {
		t := time.Now().Unix()
		value := fmt.Sprintf("SESSION:%s", md5.Md5(fmt.Sprintf("%s|%s", t, req.RemoteAddr)))
		cookie.Set(res, SESSIONID, value)
		id = value
	} else {
		id = co.Value
	}
	return id
}

func GetSession(res http.ResponseWriter, req *http.Request) map[string]interface{} {
	id := GetSessionID(res, req)

	r := redis.Get()
	defer r.Close()
	data, _ := r.Get(id)

	s := map[string]interface{}{}
	if i := other.To([]byte(data)); i != nil {
		switch i.(type) {
		case map[string]interface{}:
			s = i.(map[string]interface{})
		}

	}
	return s
}

func SaveSession(res http.ResponseWriter, req *http.Request, data map[string]interface{}) {
	id := GetSessionID(res, req)

	if len(data) <= 0 {
		return
	}

	r := redis.Get()
	defer r.Close()
	r.Set(id, other.ToByte(data))
	r.Expire(id, SESSION_TIME)
}
