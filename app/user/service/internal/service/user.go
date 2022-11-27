package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-mono-repo/api/user/service/v1"
	"kratos-mono-repo/app/user/service/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server"))}
}

func (s *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (*v1.GetUserInfoReply, error) {
	s.log.WithContext(ctx).Infof("GetUserInfo Received: %v", req)
	rv, err := s.uc.GetUserInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserInfoReply{
		Id:       rv.ID,
		Username: rv.Username,
	}, nil
}
