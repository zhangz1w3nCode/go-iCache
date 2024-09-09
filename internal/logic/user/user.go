package user

import (
	"context"
	"visual-state-machine/internal/entity"
	"visual-state-machine/internal/repo/user"
)

type User struct {
	api user.API
}

func New() *User {
	return &User{
		api: user.New(),
	}
}

func (u *User) GetUser(ctx context.Context, ID int64) (*entity.User, error) {

	param := &entity.User{
		ID:       ID,
		UserName: "zzw",
	}
	userDB, err := u.api.GetUser(ctx, param)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}
