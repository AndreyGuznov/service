package cache

import (
	"serv/pkg/logger"
	"sync"

	"github.com/astaxie/beego/cache"
)

var (
	Instance cache.Cache
	initC    sync.Once
	timeToGC string = "{'interval':1660}"
)

func InitCache() {
	initC.Do(func() {
		var err error
		Instance, err = cache.NewCache("memory", timeToGC)
		if err != nil {
			logger.Err("Cache not initialized", err)
		}
	})
}
