package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	ListArticles(ctx context.Context, opts ...ListOption) ([]*Article, error)
	FeedArticles(ctx context.Context, opts ...ListOption) ([]*Article, error)
	GetArticle(ctx context.Context, slug string) (*Article, error)
	CreateArticle(ctx context.Context, a *Article) (*Article, error)
	UpdateArticle(ctx context.Context, a *Article) (*Article, error)
	DeleteArticle(ctx context.Context, slug string) (*Article, error)

	FavoriteArticle(ctx context.Context, slug string) (*Article, error)
	UnfavoriteArticle(ctx context.Context, slug string) (*Article, error)
}

type CommentRepo interface {
}

type TagRepo interface {
	GetTags(ctx context.Context) ([]*Tag, error)
}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	tr TagRepo
	pr ProfileRepo

	log *log.Helper
}

type Article struct {
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Author      *User
}

type Tag string

func NewSocialUsecase(ar ArticleRepo,
	pr ProfileRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, pr: pr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) GetProfile(ctx context.Context, username string) (rv *Profile, error error) {
	return uc.pr.GetProfile(ctx, username)
}

func (uc *SocialUsecase) FollowUser(ctx context.Context, username string) (rv *Profile, error error) {
	return uc.pr.FollowUser(ctx, username)
}

func (uc *SocialUsecase) UnfollowUser(ctx context.Context, username string) (rv *Profile, error error) {
	return uc.pr.UnfollowUser(ctx, username)
}

func (uc *SocialUsecase) GetArticle(ctx context.Context, slug string) (rv *Article, error error) {
	return uc.ar.GetArticle(ctx, slug)
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context, in *Article) (rv *Article, error error) {
	return uc.ar.CreateArticle(ctx, in)
}

func (uc *SocialUsecase) DeleteArticle(ctx context.Context, slug string) (rv *Article, error error) {
	return uc.ar.DeleteArticle(ctx, slug)
}

func (uc *SocialUsecase) AddComment(ctx context.Context, in *Article) (rv *Article, error error) {
	return uc.ar.DeleteArticle(ctx, slug)
}

func (uc *SocialUsecase) GetComments(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) DeleteComment(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) FeedArticles(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) ListArticles(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) UpdateArticle(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) GetTags(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) FavoriteArticle(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}

func (uc *SocialUsecase) UnfavoriteArticle(ctx context.Context, in *Article) (rv *Article, error error) {
	return nil, nil
}
