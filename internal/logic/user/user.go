package user

import (
	"context"
	"fmt"
	"strconv"
	"visual-state-machine/internal/entity/model"
	"visual-state-machine/internal/repo/user"
	"visual-state-machine/internal/utils/cache"
)

type User struct {
	api             user.API
	userCacheHolder *cache.UserCacheHolder
}

func NewUserLogic(api user.API) *User {
	return &User{
		api:             api,
		userCacheHolder: cache.NewRistrettoCacheHolder(),
	}
}

func (u *User) GetUser(ctx context.Context, id string) (*model.User, error) {

	fmt.Println(u.userCacheHolder.RisCache.GetCacheStatus())

	// 获取数据
	userCache := u.userCacheHolder.RisCache.Get("user_id_" + id)
	if userCache != nil {
		fmt.Println("cache hit")
		return userCache.Data.(*model.User), nil
	}

	//查询数据库
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	userDB, err := u.api.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	//存入goCache
	idStr := strconv.FormatInt(userDB.ID, 10)
	u.userCacheHolder.RisCache.Set("user_id_"+idStr, userDB)
	fmt.Println("db hit")

	return userDB, nil
}

func (u *User) List(ctx context.Context) ([]*model.User, error) {

	users, err := u.api.List(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
