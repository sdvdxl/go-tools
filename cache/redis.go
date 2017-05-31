package cache

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"encoding/gob"

	"github.com/garyburd/redigo/redis"
	// goCache "github.com/patrickmn/go-cache"
	"github.com/sdvdxl/go-tools/errors"

	log "github.com/Sirupsen/logrus"
)

type Redis struct {
	//cacheStore *cache.Cache
	pool *redis.Pool
	// cacheStore *goCache.Cache
	// Host redis 主机
	Host string
	// Port 端口
	Port int
	// Password 密码，如果没有密码可以不设置
	Password    string
	MaxIdle     int
	IdleTimeout time.Duration
}

// Init 初始化
func (r *Redis) Init() {
	if r.MaxIdle <= 0 {
		r.MaxIdle = 2
	}

	if r.IdleTimeout <= 0 {
		r.IdleTimeout = 120
	}

	log.Info("init cache ...")
	pool := &redis.Pool{
		MaxIdle:     r.MaxIdle,
		IdleTimeout: r.IdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%v:%v", r.Host, r.Port))
			errors.Panic(err)
			if r.Password != "" {
				if _, err := c.Do("AUTH", r.Password); err != nil {
					errors.Panic(err)
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			errors.Panic(err)
			return err
		},
	}
	_, err := pool.Dial()
	if err != nil {
		log.WithError(err).WithField("host", r.Host).WithField("port", r.Port).Error("redis connect error")
		errors.Panic(err)
	}

	r.pool = pool
}

// Get 根据key获取值
// 如果存在key，则返回值，true，如果不存在，返回nil，false
func (r *Redis) Get(key string, target interface{}) bool {
	if "" == key {
		return false
	}

	value := reflect.ValueOf(target)
	if value.Type().Kind() != reflect.Ptr {
		errors.Panic(errors.ConstError("target should be pointer"))
	}
	if !r.Exists(key) {
		return false
	}

	conn := r.pool.Get()
	defer conn.Close()
	if v, err := conn.Do("GET", key); err != nil {
		panic(err)
	} else {
		inBuf := bytes.NewBuffer(v.([]byte))
		if err := gob.NewDecoder(inBuf).Decode(target); err != nil {
			panic(err)
		} else {
			log.WithFields(log.Fields{"data": target, "key": key}).Debug("get value from cache")
			return true
		}
	}
}

// GetConn 获取链接
func (r Redis) GetConn() redis.Conn {
	return r.pool.Get()
}

// Set 放置值
func (r Redis) Set(key string, value interface{}) {
	r.SetExpired(key, value, 0)
}

func (r Redis) GetString(key string) string {
	var value string
	r.Get(key, &value)
	return value
}

func (r Redis) GetFloat64(key string) float64 {
	var value float64
	r.Get(key, &value)
	return value
}

func (r Redis) GetBool(key string) bool {
	var value bool
	r.Get(key, &value)
	return value
}

func (r Redis) GetInt(key string) int {
	var value int
	r.Get(key, &value)
	return value
}

// SetExpired 放置值
// 如果value 为非 golang 基本类型，则转换成,
// 如果 expire 为<=0，那么不设置过期时间
func (r Redis) SetExpired(key string, value interface{}, expire time.Duration) {
	// log.DebugWithFields(log.Fields{"key": key, "value": value, "expireTime": expire}, "set to cache")

	var buf bytes.Buffer // Stand-in for a network connection
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		panic(err)
	}
	conn := r.GetConn()
	defer conn.Close()

	if int64(expire.Seconds()) > int64(0) {
		_, err := conn.Do("SET", key, buf.Bytes(), "EX", int64(expire.Seconds()))
		errors.Panic(err)
	} else {
		_, err := conn.Do("SET", key, buf.Bytes())
		errors.Panic(err)
	}
}

// Exists 判断key是否存在
func (r Redis) Exists(key string) bool {
	conn := r.GetConn()
	defer conn.Close()
	if exists, err := redis.Bool(conn.Do("EXISTS", key)); err != nil {
		panic(err)
	} else {
		return exists
	}
}

// Delete 删除key
func (r Redis) Delete(key string) {
	conn := r.GetConn()
	defer conn.Close()

	if r, err := conn.Do("DEL", key); err != nil {
		panic(err)
	} else {
		_ = r
		log.WithFields(log.Fields{"reply": r, "key": key}).Debug("delete from cache")
	}
}

// Close 关闭redis 链接
func (r Redis) Close() {
	if r.pool == nil {
		log.Info("redis pool already closed")
		return
	}
	if err := r.pool.Close(); err != nil {
		panic(err)
	}

	log.Info("redis pool closed")
}
