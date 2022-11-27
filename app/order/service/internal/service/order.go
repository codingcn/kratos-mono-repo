package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-mono-repo/api/order/service/v1"
	"kratos-mono-repo/app/order/service/internal/biz"
)

type OrderService struct {
	pb.UnimplementedOrderServer

	oc  *biz.OrderUsecase
	log *log.Helper
}

func NewOrderService(uc *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server")),
	}
}

func (s *OrderService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Order Hello Received: %v", req)
	s.oc.Hello(ctx, req.Id)
	return &pb.HelloReply{}, nil
}
