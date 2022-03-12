package data

import (
	"context"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Article Article
	Body    string
	Author  User
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

func (r *commentRepo) AddComment(ctx context.Context, in *biz.Comment) (rv *biz.Comment, err error) {
	c := Comment{
		Article: Article{},
		Body:    in.Body,
		Author: User{},
	}
	result := r.data.db.Create(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.Comment{
		Article: &biz.Article{},
		Body:    c.Body,
		Author: &biz.Profile{
			Username: c.Author.Username,
			Bio:      c.Author.Bio,
			Image:    c.Author.Image,
		},
	}, nil
}

func (r *commentRepo) ListComments(ctx context.Context, slug string) (rv []*biz.Comment, err error) {
	var comments []biz.Comment
	result := r.data.db.Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	rv = make([]*biz.Comment, len(comments))
	for _, x := range comments {
		rv = append(rv, &biz.Comment{
			Article: nil, // fixme
			Body:    x.Body,
			Author: &biz.Profile{
				Username: x.Author.Username,
				Bio:      x.Author.Bio,
				Image:    x.Author.Image,
			},
		})
	}
	return rv, result.Error
}

func (r *commentRepo) GetComment(ctx context.Context, id uint) (*biz.Comment, error) {
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

func (r *commentRepo) DeleteComment(ctx context.Context, id uint) (err error) {
	return r.data.db.Delete(&Comment{}, id).Error
}
