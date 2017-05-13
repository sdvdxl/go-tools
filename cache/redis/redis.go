package redis

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"encoding/gob"

	"github.com/garyburd/redigo/redis"
	// goCache "github.com/patrickmn/go-cache"

	log "github.com/Sirupsen/logrus"
	"uke.cloud/travel/util/config"
)

var (
	//cacheStore *cache.Cache
	pool *redis.Pool
	// cacheStore *goCache.Cache
	once sync.Once
	// Host redis 主机
	Host string
	// Port 端口
	Port int
	// Password 密码，如果没有密码可以不设置
	Password string
)

type redisData struct {
	Data interface{}
}

// Init 初始化
func Init() {
	once.Do(func() {
		config.Init()

		log.Info("init cache ...")
		pool = newPool(fmt.Sprintf("%v:%v", Host, Port), Password)

		if conn, err := pool.Dial(); err != nil {
			log.WithError(err).WithField("host", config.RedisHost).WithField("port", config.RedisPort).Error("redis config error")
			panic(err)
		} else { // 先放着，这里redis出错，不会使用内存
			conn.Close()
			// cacheStore = goCache.New(5*time.Minute, 30*time.Second)
			// log.Info("cache inited")
		}
	})
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

// Get 根据key获取值
// 如果存在key，则返回值，true，如果不存在，返回nil，false
func Get(key string) (interface{}, bool) {

	if !Exists(key) {
		return nil, false
	}

	conn := pool.Get()
	defer conn.Close()
	if v, err := conn.Do("GET", key); err != nil {
		panic(err)
	} else {
		inBuf := bytes.NewBuffer(v.([]byte))
		decodeValue := &redisData{}
		if err := gob.NewDecoder(inBuf).Decode(decodeValue); err != nil {
			panic(err)
		} else {
			// log.DebugWithFields(log.Fields{"data": decodeValue.Data, "key": key}, "get value from cache")
			return decodeValue.Data, true
		}
	}
}

// Set 放置值
// 如果value 为非 golang 基本类型，则转换成,
// 如果 expire 为<=0，那么不设置过期时间
func Set(key string, value interface{}, expire time.Duration) {
	// log.DebugWithFields(log.Fields{"key": key, "value": value, "expireTime": expire}, "set to cache")

	var buf bytes.Buffer // Stand-in for a network connection
	if err := gob.NewEncoder(&buf).Encode(redisData{value}); err != nil {
		panic(err)
	}
	conn := pool.Get()
	defer conn.Close()
	conn.Do("SET", key, buf.Bytes())
	if int64(expire.Seconds()) > int64(0) {
		conn.Do("EXPIRE", key, int64(expire.Seconds()))
	}

	// Flush()
	// cacheStore.Set(key, value, expire)
}

// Exists 判断key是否存在
func Exists(key string) bool {
	conn := pool.Get()
	defer conn.Close()
	if exists, err := redis.Bool(conn.Do("EXISTS", key)); err != nil {
		panic(err)
	} else {
		return exists
	}
}

// Delete 删除key
func Delete(key string) {
	conn := pool.Get()
	defer conn.Close()
	// log.DebugWithFields(log.Fields{"key": key}, "delete from cache")
	if r, err := conn.Do("DEL", key); err != nil {
		panic(err)
	} else {
		_ = r
		// log.DebugWithFields(log.Fields{"reply": r, "key": key}, "delete from cache")
	}

	// Flush()
}

// func Add(key string, value interface{}, expire time.Duration) error {
// 	// return cacheStore.Add(key, value, expire)
// }
// func Replace(key string, data interface{}, expire time.Duration) error {
// 	// return cacheStore.Replace(key, data, expire)
// }

// func Increment(key string, data int64) error {
// 	// return cacheStore.Increment(key, data)
// }
// func Decrement(key string, data int64) error {
// 	// return cacheStore.Decrement(key, data)
// }

// Flush flush数据到redis server

// Close 关闭redis 链接
func Close() {
	pool.Close()
	log.Info("redis pool closed")
}
