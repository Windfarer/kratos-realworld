package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Article struct {
	gorm.Model
	Slug        string `gorm:"size:200"`
	Title       string `gorm:"size:200"`
	Description string `gorm:"size:200"`
	Body        string
	Tags        []Tag `gorm:"many2many:article_tags;"`
	AuthorID    uint
	Author      User
}

type Tag struct {
	gorm.Model
	Name     string    `gorm:"size:200;uniqueIndex"`
	Articles []Article `gorm:"many2many:article_tags;"`
}

type Following struct {
	gorm.Model
	UserID      uint
	User        User
	FollowingID uint
	Following   User
}

type ArticleFavorite struct {
	gorm.Model
	UserID    uint
	ArticleSlug string
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func convertArticle(x Article) *biz.Article {
	return &biz.Article{
		Slug:        x.Slug,
		Title:       x.Title,
		Description: x.Description,
		Body:        x.Body,
		CreatedAt:   x.CreatedAt,
		UpdatedAt:   x.UpdatedAt,
		Author: &biz.Profile{
			Username: x.Author.Username,
			Bio:      x.Author.Bio,
			Image:    x.Author.Image,
		},
	}
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) List(ctx context.Context, opts ...biz.ListOption) (rv []*biz.Article, err error) {

	var articles []Article
	result := r.data.db.Preload("Author").Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	spew.Dump(articles)
	rv = make([]*biz.Article, len(articles))
	for i, x := range articles {
		rv[i] = convertArticle(x)
	}
	return rv, nil
}

func (r *articleRepo) Get(ctx context.Context, slug string) (rv *biz.Article, err error) {
	x := new(Article)
	err = r.data.db.Where("slug = ?", slug).First(&x).Error
	if err != nil {
		return nil, err
	}
	var fc int64
	err = r.data.db.Where("article_slug = ?", slug).Count(&fc).Error

	return &biz.Article{
		Slug:           x.Slug,
		Title:          x.Title,
		Body:           x.Body,
		Description:    x.Description,
		CreatedAt:      x.CreatedAt,
		UpdatedAt:      x.UpdatedAt,
		FavoritesCount: uint32(fc),
		Author: &biz.Profile{
			Username: x.Author.Username,
			Bio:      x.Author.Bio,
			Image:    x.Author.Image,
		},
	}, nil
}

func (r *articleRepo) Create(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	tags := make([]Tag, 0)
	for _, x := range a.TagList {
		tags = append(tags, Tag{
			Name: x,
		})
	}
	if len(tags) > 0 {
		err := r.data.db.Clauses(clause.OnConflict{DoNothing: true}).Create(tags).Error
		if err != nil {
			return nil, err
		}
	}

	po := Article{
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		Author:      User{Model:gorm.Model{ID: a.AuthorUserID}},
		Tags:        tags,
	}
	result := r.data.db.Create(&po)
	if result.Error != nil {
		return nil, result.Error
	}

	return convertArticle(po), nil
}

func (r *articleRepo) Update(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	var po Article
	if result := r.data.db.First(&po); result.Error != nil {
		return nil, result.Error
	}
	err := r.data.db.Model(&po).Updates(a).Error
	return convertArticle(po), err
}

func (r *articleRepo) Delete(ctx context.Context, slug string) error {
	rv := r.data.db.Delete(&Article{}, slug)
	return rv.Error
}

func (r *articleRepo) Favorite(ctx context.Context, currentUserID uint, slug string) error {
	af := ArticleFavorite{
		UserID:    currentUserID,
		ArticleSlug: slug,
	}
	return r.data.db.Create(&af).Error
}

func (r *articleRepo) Unfavorite(ctx context.Context, currentUserID uint, slug string) error {
	po := ArticleFavorite{
		UserID:    currentUserID,
		ArticleSlug: slug,
	}
	return r.data.db.Delete(&po).Error
}

func (r *articleRepo) GetFavoritesStatus(ctx context.Context, currentUserID uint, slugs []string) (favorited []bool, err error) {
	var po ArticleFavorite
	if result := r.data.db.First(&po); result.Error != nil {
		return nil, nil
	}
	return nil, nil
}

func (r *articleRepo) GetFavoritesCounts(ctx context.Context, slugs []string) (count []uint32, err error) {
	var po ArticleFavorite
	if result := r.data.db.First(&po); result.Error != nil {
		return nil, nil
	}
	return nil, nil
}

func (r *articleRepo) ListTags(ctx context.Context) (rv []biz.Tag, err error) {
	var tags []Tag
	err = r.data.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	rv = make([]biz.Tag, len(tags))
	for i, x := range tags {
		rv[i] = biz.Tag(x.Name)
	}
	return rv, nil
}
