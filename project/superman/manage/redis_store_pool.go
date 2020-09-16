package manage

import (
	"context"
	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/common/log"
	"shuai/superman/model"
	"time"
)

var (
	pool        *redis.Pool
	redisServer = flag.String("redisServer", ":6379", "")
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{

		MaxIdle:     4,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func InitRedisPool() {
	log.Infof("连接池地址=%v", *redisServer)
	pool = newPool(*redisServer)
}

func Get(ctx context.Context) {
	conn := pool.Get()
	rs, err := conn.Do("SET", "shuai", "124003")
	log.Infof("rs = %+v,err=%+v", rs, err)
	res, _ := conn.Do("GET", "shuai", "123", "", "")
	println(res)
	conn.Close()
}

func exec(cmds []model.Command) {
	conn := pool.Get()

	for _, cmd := range cmds {
		conn.Do(cmd.Name, cmd.Key, cmd.Value, cmd.ExpireTime)
	}
}
