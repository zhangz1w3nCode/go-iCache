package user

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"visual-state-machine/config"
	"visual-state-machine/internal/entity/model"
)

type API interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

type impl struct {
	db *gorm.DB
}

func New() API {
	db, err := gorm.Open(mysql.Open(config.Get().Database.DataSourceName), &gorm.Config{
		SkipDefaultTransaction: config.Get().Database.GormCfg.SkipDefaultTransaction,
	})
	if err != nil {
		panic(err)
	}
	return &impl{
		db: db,
	}
}
