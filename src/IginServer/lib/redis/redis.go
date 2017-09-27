package redis

import (
	// "fmt"
	"IginServer/conf"
	"github.com/garyburd/redigo/redis"
	// "net/http"
	"strconv"
)

// 重写生成连接池方法
func newPool() *redis.Pool {
	REDIS_MAX_IDLE, _ := strconv.Atoi(conf.GET["config"]["REDIS_MAX_IDLE"])
	REDIS_MAX_ACTIVE, _ := strconv.Atoi(conf.GET["config"]["REDIS_MAX_ACTIVE"])
	REDIS_NETWORK, _ := conf.GET["config"]["REDIS_NETWORK"]
	REDIS_ADDRESS, _ := conf.GET["config"]["REDIS_ADDRESS"]
	return &redis.Pool{
		MaxIdle:   REDIS_MAX_IDLE,
		MaxActive: REDIS_MAX_ACTIVE, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(REDIS_NETWORK, REDIS_ADDRESS)
			c.Do("AUTH", conf.GetString("config", "REDIS_PASSWD"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func Get() *Redis {
	p := pool.Get()
	if p.Err() != nil {
		p.Close()
		pool.Close()
		pool = newPool()
		p = pool.Get()
	}
	return &Redis{p}
}

var pool *redis.Pool

func init() {
	pool = newPool()
}
