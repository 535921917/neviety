package manage

import (
	"code.byted.org/gopkg/logs"
	"context"
	"time"
)

func CmdsExec(cmdName, key string, value interface{}, args ...interface{}) {
	conn := pool.Get()

	_, _ = conn.Do(cmdName, "shuai", "1243")

}

func Set(ctx context.Context, key string) {
	RedisClient.Set(ctx, key, "test", time.Minute)
	v, _ := RedisClient.Get(ctx, key).Result()
	logs.CtxInfo(ctx, v)
}

