package user

import (
	"context"
	"strconv"

	"visual-state-machine/internal/entity/model"
	"visual-state-machine/internal/repo/user"
)

type User struct {
	api user.API
}

func New(api user.API) *User {
	return &User{
		api: api,
	}
}

func (u *User) GetUser(ctx context.Context, id string) (*model.User, error) {
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	param := &model.User{
		ID: ID,
	}
	userDB, err := u.api.GetUser(ctx, param)

	if err != nil {
		return nil, err
	}

	return userDB, nil
}
