package redis_conn

import (
	"High-concurrent-spike-system/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	RedisPool *redis.Pool
)

func init() {
	RedisPool = newRedisPool()
}

// 创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.MyConfig.Redis.Maxidle,
		MaxActive:   config.MyConfig.Redis.MaxActive,
		IdleTimeout: config.MyConfig.Redis.IdleTimeout * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			dial, e := redis.Dial("tcp", config.MyConfig.Redis.Port)
			if e != nil {
				log.Println(e.Error())
				panic(e.Error())
			}
			return dial, e
		},
		// 检测连接状况
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

