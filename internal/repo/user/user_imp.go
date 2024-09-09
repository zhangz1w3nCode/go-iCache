package user

import (
	"context"
	"visual-state-machine/internal/entity"
)

func (i *impl) GetUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	userDB := &entity.User{}
	if err := i.db.WithContext(ctx).Model(&entity.User{}).
		Where("id = ?", user.ID).
		Find(userDB).
		Error; err != nil {
		return nil, err
	}
	return userDB, nil
}
