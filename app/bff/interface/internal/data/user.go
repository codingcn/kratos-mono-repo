package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	userv1 "kratos-mono-repo/api/user/service/v1"
	"kratos-mono-repo/app/bff/interface/internal/biz"
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
	//fmt.Println("===========", telnet.TcpGather("consul", []string{"8500"}))
	//fmt.Println("===========", telnet.TcpGather("consul", []string{"8500"}))
	if err != nil {
		r.log.WithContext(ctx).Error("GetUserInfo RPC error", zap.Error(err))
		return nil, err
	}
	r.log.WithContext(ctx).Infof("GetUserInfo RPC success")
	return &biz.User{
		Id:       reply.Id,
		Username: reply.Username,
	}, nil
}
