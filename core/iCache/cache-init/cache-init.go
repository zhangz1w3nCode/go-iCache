package cacheInit

import (
	"github.com/spf13/viper"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-manager"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type CacheInit struct {
	CacheManager   *cacheManager.CacheManager
	GormConnection *gorm.DB
}

func NewCacheInit() *CacheInit {
	//初始化数据库连接
	gormConnection, err := gorm.Open(mysql.Open(viper.GetString("cache-config.database.data_source_name")), &gorm.Config{
		SkipDefaultTransaction: viper.GetBool("cache-config.database.gorm_cfg.skip_default_transaction"),
	})
	if err != nil {
		log.Fatalf("Init database error:%v", err)
	}
	sqlDB, _ := gormConnection.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("cache-config.database.connection_max_idle_connections")) // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("cache-config.database.max_open_connections"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("cache-config.database.connection_max_lifetime")) * time.Millisecond) // 设置打开数据库连接的最大数量 // 设置连接可复用的最大时

	//初始化cacheManager
	cacheManager := cacheManager.NewCacheManager()
	return &CacheInit{
		CacheManager:   cacheManager,
		GormConnection: gormConnection,
	}
}
