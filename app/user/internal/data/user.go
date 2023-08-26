package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	orderv1 "kratos-mono-repo/api/order/v1"
	"kratos-mono-repo/app/user/internal/biz"
	"kratos-mono-repo/app/user/internal/pkg/util"
	"time"
)

//var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

var userCacheKey = func(username string) string {
	return "user_cache_key_" + username
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = ph
	err = r.data.db.WithContext(ctx).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (r *userRepo) GetUserInfo(ctx context.Context, id uint64) (*biz.User, error) {
	r.data.oc.Hello(ctx, &orderv1.HelloRequest{
		Id: id,
	})
	return &biz.User{
		ID:       id,
		Username: "用户微服务",
		Nickname: "",
		Password: "",
		Phone:    "",
		Gender:   0,
		Email:    "",
		Avatar:   "",
	}, nil
}
func (r *userRepo) GetUserInfo2(ctx context.Context, id uint64) (*biz.User, error) {
	// try to fetch from cache
	cacheKey := userCacheKey(fmt.Sprintf("%d", id))
	target, err := r.getUserFromCache(ctx, cacheKey)
	if err != nil {
		// fetch from db while cache missed
		u := biz.User{}
		err = r.data.db.First(&u).Error
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		// set cache
		r.setUserCache(ctx, target, cacheKey)
	}
	return target, nil
}

func (r *userRepo) getUserFromCache(ctx context.Context, key string) (*biz.User, error) {
	result, err := r.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheUser = &biz.User{}
	err = json.Unmarshal([]byte(result), cacheUser)
	if err != nil {
		return nil, err
	}
	return cacheUser, nil
}

func (r *userRepo) setUserCache(ctx context.Context, user *biz.User, key string) {
	marshal, err := json.Marshal(user)
	if err != nil {
		r.log.Errorf("fail to set server cache:json.Marshal(%v) error(%v)", user, err)
	}
	err = r.data.redisCli.Set(ctx, key, string(marshal), time.Minute*30).Err()
	if err != nil {
		r.log.Errorf("fail to set server cache:redis.Set(%v) error(%v)", user, err)
	}
}
