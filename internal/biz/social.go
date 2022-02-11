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
	CreateArticle(ctx context.Context, a Article) (*Article, error)
	UpdateArticle(ctx context.Context, a Article) (*Article, error)
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
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}
