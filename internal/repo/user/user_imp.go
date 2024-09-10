package user

import (
	"context"

	"visual-state-machine/internal/entity/model"
)

func (i *impl) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
	userDB := &model.User{}
	if err := i.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", user.ID).
		Find(userDB).
		Error; err != nil {
		return nil, err
	}
	return userDB, nil
}
