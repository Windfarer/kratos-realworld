package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ArticleSlug string
	Article     Article `gorm:"references:Slug"`
	Body        string
	AuthorID    uint
	Author      User
}

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentRepo) Create(ctx context.Context, in *biz.Comment) (rv *biz.Comment, err error) {
	c := Comment{
		ArticleSlug: in.Article.Slug,
		Body:    in.Body,
		AuthorID: in.AuthorID,
	}
	result := r.data.db.Create(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Comment{
		ID: c.ID,
		Article: &biz.Article{},
		Body:    c.Body,
		Author: &biz.Profile{
			Username: c.Author.Username,
			Bio:      c.Author.Bio,
			Image:    c.Author.Image,
		},
	}, nil
}

func (r *commentRepo) List(ctx context.Context, slug string) (rv []*biz.Comment, err error) {
	var comments []Comment
	result := r.data.db.Where("article_slug = ?", slug).Preload("Author").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	rv = make([]*biz.Comment, len(comments))
	for i, x := range comments {
		rv[i] = &biz.Comment{
			ID: x.ID,
			Article: nil, // fixme
			Body:    x.Body,
			Author: &biz.Profile{
				Username: x.Author.Username,
				Bio:      x.Author.Bio,
				Image:    x.Author.Image,
			},
		}
	}
	return rv, result.Error
}

func (r *commentRepo) Get(ctx context.Context, id uint) (*biz.Comment, error) {
	var c Comment
	result := r.data.db.First(&c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Comment{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Body:      c.Body,
		Author: &biz.Profile{
			Username: c.Author.Username,
			Bio:      c.Author.Bio,
			Image:    c.Author.Image,
		},
	}, nil
}

func (r *commentRepo) Delete(ctx context.Context, id uint) (err error) {
	return r.data.db.Delete(&Comment{}, id).Error
}
