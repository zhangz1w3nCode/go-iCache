package cacheStorage

import (
	"encoding/json"
	cacheInit "github.com/zhangz1w3nCode/go-iCache/core/iCache/cache-init"
	"github.com/zhangz1w3nCode/go-iCache/core/iCache/cache/go-cache"
	"github.com/zhangz1w3nCode/go-iCache/internal/entity/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var logger *zap.Logger

func CacheStorage(cacheInit *cacheInit.CacheInit, serviceName, serviceAddress string) {
	//初始化日志 logger
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	//日志：将缓存的每个具体的缓存 批量快速的写入到数据库中持久化
	logger.Info("Start storage cache data to database!")
	//开始统计耗时
	startTime := time.Now()
	//1.获取mysql连接 dbConn
	db := cacheInit.GormConnection
	//2.从获取managerCache中该服务每个缓存的所有key和value信息
	cacheDetail := cacheInit.CacheManager.GetCacheDetail()
	//拿到所有的缓存
	var cacheList []*model.CacheErrorRecover
	for cacheName := range cacheDetail {
		cache := cacheDetail[cacheName].(*go_cache.GoCache)
		keys := cache.GetKeys()
		for _, key := range keys {
			wrapper := cache.Get(key)
			value := wrapper.Data
			valueStr, _ := json.Marshal(value)
			cacheData := &model.CacheErrorRecover{
				CacheName:      cacheName,
				CacheKey:       key,
				CacheValue:     string(valueStr),
				CacheType:      "go_cache",
				ErrorType:      "ServerError",
				ErrorMessage:   "Server is error closing!",
				ServiceName:    serviceName,
				ServiceAddress: serviceAddress,
				IsRecover:      false,
				CreateTime:     time.Now(),
				UpdateTime:     time.Now(),
				CreateBy:       "admin",
				UpdateBy:       "admin",
			}
			cacheList = append(cacheList, cacheData)
		}
	}
	//3.通过dbConn将数据存入MySQL
	semaphoreSize := 10
	batchSize := 5000
	semaphore := make(chan struct{}, semaphoreSize)
	waitGroup := sync.WaitGroup{}
	storage(db, cacheList, batchSize, &waitGroup, semaphore)
	waitGroup.Wait()
	//4.日志：将缓存的每个具体的缓存 批量快速的写入到数据库中持久化
	logger.Info("End storage cache data to database!")
	//结束统计耗时
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	logger.Info("Storage cache data to database cost time:", zap.Duration("elapsedTime", elapsedTime))
}

func storage(db *gorm.DB, cacheList []*model.CacheErrorRecover, batchSize int, waitGroup *sync.WaitGroup, semaphore chan struct{}) {
	for i := 0; i < len(cacheList); i += batchSize {
		semaphore <- struct{}{}
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			defer func() {
				<-semaphore
			}()
			end := i + batchSize
			if end > len(cacheList) {
				end = len(cacheList)
			}
			dataList := cacheList[i:end]
			if len(dataList) == 0 {
				// 没有数据需要处理
				log.Fatalf("No data to process.....")
			}
			err := db.Omit("recover_time", "delete_time", "delete_by").CreateInBatches(dataList, len(dataList)).Error
			if err != nil {
				log.Fatalf("CreateInBatches error: %v", err)
			}
		}()
	}
}
