package start

import (
	"github.com/spf13/viper"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/manager"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type CacheInit struct {
	CacheManager *manager.CacheManager
	db           *gorm.DB
}

func NewCacheInit() *CacheInit {
	db, err := gorm.Open(mysql.Open(viper.GetString("config.database.data_source_name")), &gorm.Config{
		SkipDefaultTransaction: viper.GetBool("config.database.gorm_cfg.skip_default_transaction"),
	})
	if err != nil {
		log.Fatalf("Init database error:%v", err)
	}
	return &CacheInit{
		CacheManager: manager.NewCacheManager(),
		db:           db,
	}
}
