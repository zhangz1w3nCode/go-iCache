package calling

import (
	"context"
	"visual-state-machine/internal/entity"
	"visual-state-machine/internal/repo/log"
)

type UserLogic struct {
	api log.API
}

func New() *UserLogic {
	return &UserLogic{
		api: log.New(),
	}
}

func (l *UserLogic) GetUser(ctx context.Context, ID int64) (*entity.OceanUser, error) {

	//构建callingLog的更新参数
	param := &entity.OceanUser{
		ID:       ID,
		UserName: "zzw",
	}
	userDB, err := l.api.GetUser(ctx, param)
	//userDB, err := log.New().GetUser(ctx, param)
	if err != nil {
		return nil, err
	}

	return userDB, nil
}
