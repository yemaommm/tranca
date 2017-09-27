package redis

import (
	"github.com/garyburd/redigo/redis"
)

type PubSubConn struct {
	redis.PubSubConn
}

func PUB(conn redis.Conn) *PubSubConn {
	return &PubSubConn{redis.PubSubConn{conn}}
}

type Redis struct {
	redis.Conn
}

func (r *Redis) TTL(key interface{}) (int, error) {
	return redis.Int(r.Do("TTL", key))
}

func (r *Redis) Del(key interface{}) (interface{}, error) {
	return r.Do("DEL", key)
}

func (r *Redis) Get(str interface{}) (string, error) {
	return redis.String(r.Do("GET", str))
}

func (r *Redis) MGet(str ...interface{}) ([]string, error) {
	return redis.Strings(r.Do("MGET", str...))
}

func (r *Redis) Set(key, value interface{}) (interface{}, error) {
	return r.Do("SET", key, value)
}

func (r *Redis) MSet(value ...interface{}) (interface{}, error) {
	return r.Do("MSET", value...)
}

func (r *Redis) HGet(key, value interface{}) (string, error) {
	return redis.String(r.Do("HGET", key, value))
}

func (r *Redis) HMGet(str ...interface{}) ([]string, error) {
	return redis.Strings(r.Do("HMGET", str...))
}

func (r *Redis) HGetAll(str interface{}) (map[string]string, error) {
	return redis.StringMap(r.Do("HGETALL", str))
}

func (r *Redis) HSet(table, key, value interface{}) (interface{}, error) {
	return r.Do("HSET", table, key, value)
}

func (r *Redis) HMSet(str ...interface{}) (interface{}, error) {
	return r.Do("HMSET", str...)
}

func (r *Redis) HDel(value ...interface{}) (interface{}, error) {
	return r.Do("HDEL", value...)
}

func (r *Redis) Expire(key, time interface{}) (string, error) {
	return redis.String(r.Do("EXPIRE", key, time))
}

func (r *Redis) RPUSH(key, data interface{}) (interface{}, error) {
	return r.Do("RPUSH", key, data)
}

func (r *Redis) LPUSH(key, data interface{}) (interface{}, error) {
	return r.Do("LPUSH", key, data)
}

func (r *Redis) RPOP(key interface{}) (string, error) {
	return redis.String(r.Do("RPOP", key))
}

func (r *Redis) LPOP(key interface{}) (string, error) {
	return redis.String(r.Do("LPOP", key))
}

func (r *Redis) PUB(key, value interface{}) (interface{}, error) {
	return r.Do("Publish", key, value)
}
