package manage

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"shuai/superman/consts"
	"strconv"
	"strings"
	"time"
)

func InitWatcher(ctx context.Context) {
	ctxCancel, cancel := context.WithCancel(ctx)
	w := Watcher{
		Ctx:    ctxCancel,
		Cancel: cancel,
	}
	w.StartWatcher()
}

type Watcher struct {
	Ctx    context.Context
	Cancel context.CancelFunc
	/*		stop func()
			mtx      sync.Mutex*/
}

func (w *Watcher) StartWatcher() {
	go func() {
		for {
			select {
			case <-w.Ctx.Done():
				return
			case <-w.loop():
			}
		}
	}()
}
func (w *Watcher) StopWatcher() {
	w.Cancel()
}

func (w *Watcher) loop() <-chan error {

	errChan := make(chan error)
	go func() {
		t := time.NewTicker(time.Second)
		defer func() {
			t.Stop()
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic error %v", r)
				close(errChan)
			}
		}()
		for {
			select {
			case <-t.C:
				doWork(w.Ctx)
			case <-w.Ctx.Done():
				close(errChan)
				return
			}
		}
	}()
	return errChan
}

func doWork(ctx context.Context) {
	serverNames, err := RedisClient.SMembers(ctx, consts.AllServiceNames).Result()
	log.Infof("[%v]Timer doWork check,serviceSet=%v", time.Now().Format(consts.LayoutTime), serverNames)

	if len(serverNames) == 0 || err != nil {
		return
	}
	for _, psm := range serverNames {
		psm = strings.TrimPrefix(psm, consts.ServiceNamePre)
		workOne(ctx, psm)
	}
}

func workOne(ctx context.Context, psm string) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	cnt, err := RedisClient.ZRemRangeByScore(ctx, KeyForServiceZSet(psm), "0", now).Result()
	if cnt == 0 || err != nil {
		return
	}

	pipe := RedisClient.Pipeline()
	pipe.Incr(ctx, KeyForServiceVersion(psm))
	pipe.Incr(ctx, consts.GlobalVersion)
	if RedisClient.Exists(ctx, KeyForServiceZSet(psm)).Val() == 0 {
		pipe.SRem(ctx, consts.AllServiceNames, KeyForServiceName(psm))
	}
	_, _ = pipe.Exec(ctx)

}
