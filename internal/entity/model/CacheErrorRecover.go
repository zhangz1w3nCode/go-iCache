package model

import (
	"time"
)

type CacheErrorRecover struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CacheName      string    `gorm:"column:cache_name;not null" json:"cache_name"`
	CacheKey       string    `gorm:"column:cache_key;not null" json:"cache_key"`
	CacheValue     string    `gorm:"column:cache_value;not null" json:"cache_value"`
	CacheType      string    `gorm:"column:cache_type;not null" json:"cache_type"`
	ErrorType      string    `gorm:"column:error_type;not null" json:"error_type"`
	ErrorMessage   string    `gorm:"column:error_message;not null" json:"error_message"`
	ServiceName    string    `gorm:"column:service_name;not null" json:"service_name"`
	ServiceAddress string    `gorm:"column:service_address;not null" json:"service_address"`
	IsRecover      bool      `gorm:"column:is_recover" json:"is_recover"`
	RecoverTime    time.Time `gorm:"column:recover_time" json:"recover_time"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	CreateBy       string    `gorm:"column:create_by" json:"create_by"`
	UpdateBy       string    `gorm:"column:update_by" json:"update_by"`
	DeleteTime     time.Time `gorm:"column:delete_time" json:"delete_time"`
	DeleteBy       string    `gorm:"column:delete_by" json:"delete_by"`
}

func (c *CacheErrorRecover) ToPB() {

}

func (c *CacheErrorRecover) TableName() string {
	return "cache_error_recover"
}
