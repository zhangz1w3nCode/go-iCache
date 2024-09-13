package user

import (
	"context"

	"visual-state-machine/internal/entity/model"
)

func (i *impl) Get(ctx context.Context, id int64) (*model.User, error) {
	userDB := &model.User{}
	if err := i.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		First(userDB).
		Error; err != nil {
		return nil, err
	}
	return userDB, nil
}

func (i *impl) List(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)
	if err := i.db.WithContext(ctx).Model(&model.User{}).
		Find(&users).
		Error; err != nil {
		return nil, err
	}
	return users, nil
}
