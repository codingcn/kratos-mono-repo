package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-mono-repo/api/bff/interface/v1"
	"kratos-mono-repo/app/bff/interface/internal/biz"

	pb "kratos-mono-repo/api/bff/interface/v1"
)

type UserService struct {
	v1.UnimplementedUserInterfaceServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server"))}
}

func (s *UserService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoReq) (*pb.GetUserInfoReply, error) {
	s.log.WithContext(ctx).Infof("GetUserInfo api")
	resp, err := s.uc.GetUserInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoReply{
		Code:    0,
		Message: "",
		Data: &pb.GetUserInfoReply_GetUserInfoReplyData{
			Id:       resp.Id,
			Username: resp.Username,
		},
	}, nil
}
