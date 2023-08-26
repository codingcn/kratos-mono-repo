package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	Id       uint64
	Username string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	GetUserInfo(context.Context, uint64) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserInfo(ctx context.Context, id uint64) (*User, error) {
	uc.log.WithContext(ctx).Warnf("test warn log")
	return uc.repo.GetUserInfo(ctx, id)
}
