package log

import (
	"context"
	"visual-state-machine/internal/entity"
)

func (i *impl) GetUser(ctx context.Context, user *entity.OceanUser) (*entity.OceanUser, error) {
	userDB := &entity.OceanUser{}
	if err := i.db.WithContext(ctx).Model(&entity.OceanUser{}).
		Where("id = ?", user.ID).
		Find(userDB).
		Error; err != nil {
		return nil, err
	}
	return userDB, nil
}
