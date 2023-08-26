package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v12 "kratos-mono-repo/api/bff/v1"
	"kratos-mono-repo/app/bff/internal/biz"
)

type UserService struct {
	v12.UnimplementedUserInterfaceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server"))}
}

func (s *UserService) GetUserInfo(ctx context.Context, req *v12.GetUserInfoReq) (*v12.GetUserInfoReply, error) {
	resp, err := s.uc.GetUserInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v12.GetUserInfoReply{
		Code:    0,
		Message: "",
		Data: &v12.GetUserInfoReply_GetUserInfoReplyData{
			Id:       resp.Id,
			Username: resp.Username,
		},
	}, nil
}
