package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v12 "kratos-mono-repo/api/user/v1"
	"kratos-mono-repo/app/user/internal/biz"
)

type UserService struct {
	v12.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server"))}
}

func (s *UserService) GetUserInfo(ctx context.Context, req *v12.GetUserInfoReq) (*v12.GetUserInfoReply, error) {
	s.log.WithContext(ctx).Infof("GetUserInfo Received: %v", req)
	rv, err := s.uc.GetUserInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v12.GetUserInfoReply{
		Id:       rv.ID,
		Username: rv.Username,
	}, nil
}
