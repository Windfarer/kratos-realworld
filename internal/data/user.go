package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-realworld/internal/biz"
)

type FollowUser struct {
	gorm.Model
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	return nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (*biz.User, error) {
	return nil, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *profileRepo) GetProfile(ctx context.Context, username string) (*biz.Profile, error) {
	return nil, nil
}

func (r *profileRepo) FollowUser(ctx context.Context, username string) (*biz.Profile, error) {
	return nil, nil
}

func (r *profileRepo) UnfollowUser(ctx context.Context, username string) (*biz.Profile, error) {
	return nil, nil
}
