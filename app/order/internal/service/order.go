package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mono-repo/api/order/v1"
	"kratos-mono-repo/app/order/internal/biz"
)

type OrderService struct {
	v1.UnimplementedOrderServer

	oc  *biz.OrderUsecase
	log *log.Helper
}

func NewOrderService(uc *biz.OrderUsecase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  uc,
		log: log.NewHelper(log.With(logger, "module", "server/server-server")),
	}
}

func (s *OrderService) Hello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Order Hello Received: %v", req)
	s.oc.Hello(ctx, req.Id)
	return &v1.HelloReply{}, nil
}
