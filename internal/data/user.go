package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/go-kratos/kratos/v2/errors"
)

type FollowUser struct {
	gorm.Model
	user      string
	following string
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

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
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

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (rv *biz.User, err error) {
	u := new(User)
	result := r.data.db.Where("email = ?", email).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("user", "not found by email")
	}
	if result.Error != nil {
		return nil, err
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, username string) (rv *biz.User, err error) {
	u := new(User)
	err = r.data.db.Where("username = ?", username).First(u).Error
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *profileRepo) GetProfile(ctx context.Context, username string) (rv *biz.Profile, err error) {
	u := new(User)
	err = r.data.db.Where("username = ?", username).First(u).Error
	if err != nil {
		return nil, err
	}
	return &biz.Profile{
		Username:  u.Username,
		Bio:       u.Bio,
		Image:     u.Image,
		Following: false, // fixme
	}, nil
}

func (r *profileRepo) FollowUser(ctx context.Context, currentUsername, username string) (err error) {
	po := FollowUser{
		user:      currentUsername,
		following: username,
	}
	return r.data.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&po).Error
}

func (r *profileRepo) UnfollowUser(ctx context.Context, currentUsername, username string) (err error) {
	po := FollowUser{
		user:      currentUsername,
		following: username,
	}
	return r.data.db.Delete(&po).Error
}

func (r *profileRepo) GetUserFollowingStatus(ctx context.Context, currentUsername, username string) (following bool, err error) {
	var po FollowUser
	if result := r.data.db.First(&po); result.Error != nil {
		return false, nil
	}
	return true, nil
}
