package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/cache"
)

type Memory struct {
	memory *cache.Cache
}

var (
	instance *Memory
	initMem  sync.Once
	timeToGC string = "{'interval':1660}"
)

func GetCahe() *Memory {
	initMem.Do(func() {
		var err error
		m, err := cache.NewCache("memory", timeToGC)
		if err != nil {
			fmt.Println(err)
		}
		instance = &Memory{memory: &m}
	})
	return instance
}

func (m *Memory) Put(key string, data interface{}, time time.Duration) error {
	err := cache.Cache.Put(*GetCahe().memory, key, data, time)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) Get(key string) interface{} {
	return cache.Cache.Get(*GetCahe().memory, key)
}

func (m *Memory) Delete(key string) {
	cache.Cache.Delete(*GetCahe().memory, key)
}

// var (
// 	mem      *Memory
// 	timeToGC string = "{'interval':1660}"
// )

// type Memory struct {
// 	Instance *cache.Cache
// }

// func Conn() *Memory {
// 	if mem == nil {
// 		var err error

// 		mem, err = Init()
// 		if err != nil {
// 			logger.Err("Error of initialization Cache", err)
// 			return nil
// 		}
// 		logger.Info("Cache initialized")
// 	}

// 	return mem
// }

// func doInit() (*Memory, error) {

// 	instance, err := cache.NewCache("memory", timeToGC)

// 	if err != nil {
// 		logger.Err("Failed initialization of memory cache", err)
// 		return nil, err
// 	}

// 	memor := Memory{
// 		Instance: instance,
// 	}

// 	return &memor, nil
// }

// func Init() (*Memory, error) {
// 	var memsess *Memory
// 	var err error

// 	memsess, err = doInit()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return memsess, nil
// }
