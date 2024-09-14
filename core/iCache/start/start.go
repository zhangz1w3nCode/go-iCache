package start

import (
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
)

type CacheInit struct {
	CacheManager *manager.CacheManager
}

func NewCacheInit() *CacheInit {
	return &CacheInit{
		CacheManager: manager.NewCacheManager(),
	}
}
