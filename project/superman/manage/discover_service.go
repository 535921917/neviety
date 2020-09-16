package manage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/common/log"
	"shuai/superman/consts"
	"shuai/superman/model"
	"strings"
	"time"
)

func ServiceKeep(ctx context.Context, item *model.ServiceItem) {
	now := time.Now().Unix()
	RedisClient.SAdd(ctx, consts.AllServiceNames, KeyForServiceName(item.PSM))
	cnt, _ := RedisClient.ZAdd(ctx, KeyForServiceZSet(item.PSM), &redis.Z{Score: float64(item.TTL + now), Member: item.GetKey()}).Result()
	//ZAdd == 0  <==> set 没有变化
	if cnt == 0 {
		log.Info("keep success")
		return
	}
	pipe := RedisClient.Pipeline()
	pipe.IncrBy(ctx, KeyForServiceVersion(item.PSM), 1)
	pipe.IncrBy(ctx, consts.GlobalVersion, 1)
	_, _ = pipe.Exec(ctx)
	log.Info("register success")
}

func CancelKeep(ctx context.Context, item *model.ServiceItem) {
	score, err := RedisClient.ZScore(ctx, KeyForServiceZSet(item.PSM), item.GetKey()).Result()
	log.Debugf("score=%v,err=%v", score, err)
	if score == 0 || err != nil {
		return
	}
	pipe := RedisClient.Pipeline()
	pipe.ZRem(ctx, KeyForServiceZSet(item.PSM), item.GetKey())
	pipe.IncrBy(ctx, KeyForServiceVersion(item.PSM), 1)
	pipe.IncrBy(ctx, consts.AllServiceNames, 1)
	_, _ = pipe.Exec(ctx)
}

func ServiceList(ctx context.Context, psmArray []string) interface{} {
	m := make(map[string]interface{})
	for _, psm := range psmArray {
		m[psm] = ServiceDetail(ctx, psm)
	}
	return m
}

func ServiceDetail(ctx context.Context, psm string) *model.ServerListResponse {
	serviceVersion := RedisClient.IncrBy(ctx, KeyForServiceVersion(psm), 0).Val()
	instances, _ := RedisClient.ZRange(ctx, KeyForServiceZSet(psm), 0, -1).Result()
	var items []*model.ServiceItem
	for _, instance := range instances {
		s := strings.Split(instance, ":")
		items = append(items, model.NewServiceItem(psm, s[0], s[1]))
	}
	return &model.ServerListResponse{
		Items:          items,
		ServiceVersion: serviceVersion,
		PSM:            psm,
	}
}

func ServiceVersion(ctx context.Context, psmArray []string) interface{} {
	globalVersion := RedisClient.IncrBy(ctx, consts.GlobalVersion, 0).Val()

	versionMap := make(map[string]interface{})
	for _, psm := range psmArray {
		serviceVersion := RedisClient.IncrBy(ctx, KeyForServiceVersion(psm), 0).Val()
		versionMap[psm] = serviceVersion
	}
	return model.ServerVersionResponse{
		GlobalVersion: globalVersion,
		VersionMap:    versionMap,
	}

}
