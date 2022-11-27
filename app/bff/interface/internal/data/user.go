package data

import (
	"context"
	userv1 "kratos-mono-repo/api/user/service/v1"
	"kratos-mono-repo/app/bff/interface/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserUsecase(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserInfo(ctx context.Context, id uint64) (*biz.User, error) {
	reply, err := r.data.uc.GetUserInfo(ctx, &userv1.GetUserInfoReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, nil
}
