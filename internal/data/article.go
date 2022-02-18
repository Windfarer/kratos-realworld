package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-realworld/internal/biz"
	"time"
)

type Article struct {
	gorm.Model
	Slug        string
	Title       string
	Description string
	Body        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Tags        []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}

type Following struct {
	gorm.Model
	User      uint
	Following uint
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func (articleRepo) ListArticles(ctx context.Context, opts ...biz.ListOption) ([]*biz.Article, error) {
	return nil, nil
}

func (articleRepo) FeedArticles(ctx context.Context, opts ...biz.ListOption) ([]*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) GetArticle(ctx context.Context, slug string) (*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) CreateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) UpdateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) DeleteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) FavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	panic("implement me")
}

func (articleRepo) UnfavoriteArticle(ctx context.Context, slug string) (*biz.Article, error) {
	panic("implement me")
}

// NewGreeterRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
