package common

import (
	"cloud-disk/internal/log"

	"github.com/panjf2000/ants/v2"
)

var Pool *GoroutinePool

type GoroutinePool struct {
	goroutinePool *ants.Pool
}

func InitGoroutinePool(num int) error {
	if num <= 0 {
		return InvalidInputParamErr
	}

	pool, err := ants.NewPool(num)
	if err != nil {
		log.Error("fail to new a goroutine pool, err:%s", err)
		return err
	}
	Pool = &GoroutinePool{pool}
	return nil
}

func (g *GoroutinePool) Submit(task func()) {
	g.goroutinePool.Submit(task)
}

func Close() {
	if Pool != nil {
		Pool.goroutinePool.Release()
	}
}
