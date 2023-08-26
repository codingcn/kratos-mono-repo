package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

var (
	// ErrUserNotFound is server not found.
	//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "server not found")
	ErrUserNotFound = errors.New("server not found")
)

// Greeter is a Greeter model.
type User struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Gender   uint8  `json:"gender"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func (receiver User) TableName() string {
	return "users"
}

// UserRepo is a Greater repo.
type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUserInfo(ctx context.Context, id uint64) (*User, error)
}

// UserUseCase is a Greeter usecase.
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/server"))}
}

func (uc *UserUseCase) GetUserInfo(ctx context.Context, id uint64) (*User, error) {
	return uc.repo.GetUserInfo(ctx, id)
}
