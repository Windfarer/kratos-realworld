package biz

import (
	"context"
	"errors"
	"fmt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", b)
	return string(b)
}

func verifyPaasword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

type ProfileRepo interface {
	GetProfile(ctx context.Context, username string) (*Profile, error)
	FollowUser(ctx context.Context, username string) (*Profile, error)
	UnfollowUser(ctx context.Context, username string) (*Profile, error)
}

type UserUsecase struct {
	ur   UserRepo
	pr   ProfileRepo
	jwtc *conf.JWT

	log *log.Helper
}

type Profile struct {
}

func NewUserUsecase(ur UserRepo,
	pr ProfileRepo, logger log.Logger, jwtc *conf.JWT) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, jwtc: jwtc, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Token, username)
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    uc.generateToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPaasword(u.PasswordHash, password) {
		return nil, errors.New("login failed")
	}

	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uc.generateToken(u.Username),
	}, nil
}

func (uc *UserUsecase) GetCurrentUser(ctx context.Context) (*User, error) {
	return nil, nil
}
