package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type FollowUser struct {
	gorm.Model
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}
	rv := r.data.db.Create(&user)
	return rv.Error
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
