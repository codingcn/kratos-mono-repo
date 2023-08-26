package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Order is a Order model.
type Order struct {
	Hello string
}

// OrderRepo is a Greater repo.
type OrderRepo interface {
	Hello(context.Context, uint64) (*Order, error)
}

// OrderUsecase is a Order usecase.
type OrderUsecase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUsecase(repo OrderRepo, logger log.Logger) *OrderUsecase {
	return &OrderUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *OrderUsecase) Hello(ctx context.Context, id uint64) (*Order, error) {
	uc.log.WithContext(ctx).Infof("Order Hello: %v", id)
	return uc.repo.Hello(ctx, id)
}
