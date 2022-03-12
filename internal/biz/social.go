package biz

import (
	"context"
	"errors"
	"kratos-realworld/internal/pkg/middleware/auth"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	ListArticles(ctx context.Context, opts ...ListOption) ([]*Article, error)
	GetArticle(ctx context.Context, slug string) (*Article, error)
	CreateArticle(ctx context.Context, a *Article) (*Article, error)
	UpdateArticle(ctx context.Context, a *Article) (*Article, error)
	DeleteArticle(ctx context.Context, slug string) error

	FavoriteArticle(ctx context.Context, currentUsername string, slug string) error
	UnfavoriteArticle(ctx context.Context, currentUsername string, slug string) error
	GetArticleFavoriteStatus(ctx context.Context, currentUsername string, slug string) (favorited bool, err error)
}

type CommentRepo interface {
	AddComment(ctx context.Context, c *Comment) (*Comment, error)
	GetComment(ctx context.Context, id uint) (*Comment, error)
	ListComments(ctx context.Context, slug string) ([]*Comment, error)
	DeleteComment(ctx context.Context, id uint) error
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
	Slug           string
	Title          string
	Description    string
	Body           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TagList        []string
	Favorited      bool
	FavoritesCount uint32

	AuthorUsername       string

	Author         *Profile
}

type Comment struct {
	ID        uint
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time

	ArticleID uint
	AuthorUsername string

	Article   *Article
	Author         *Profile
}

type Tag string

func (o *Article) verifyAuthor(username string) bool {
	return o.Author.Username == username
}

func (o *Comment) verifyAuthor(username string) bool {
	return o.Author.Username == username
}

func NewSocialUsecase(ar ArticleRepo,
	pr ProfileRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, pr: pr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) GetProfile(ctx context.Context, username string) (rv *Profile, err error) {
	return uc.pr.GetProfile(ctx, username)
}

func (uc *SocialUsecase) FollowUser(ctx context.Context, username string) (rv *Profile, err error) {
	cu := auth.FromContext(ctx)
	err = uc.pr.FollowUser(ctx, cu.Username, username)
	if err != nil {
		return nil, err
	}
	rv, err = uc.pr.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *SocialUsecase) UnfollowUser(ctx context.Context, username string) (rv *Profile, err error) {
	cu := auth.FromContext(ctx)
	if err != nil {
		uc.pr.UnfollowUser(ctx, cu.Username, username)
	}
	rv, err = uc.pr.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *SocialUsecase) GetArticle(ctx context.Context, slug string) (rv *Article, err error) {
	return uc.ar.GetArticle(ctx, slug)
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context, in *Article) (rv *Article, err error) {
	u := auth.FromContext(ctx)
	in.Author.Username = u.Username
	return uc.ar.CreateArticle(ctx, in)
}

func (uc *SocialUsecase) DeleteArticle(ctx context.Context, slug string) (err error) {
	a, err := uc.ar.GetArticle(ctx, slug)
	if err != nil {
		return err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).Username) {
		return errors.New("no permission 401")
	}
	return uc.ar.DeleteArticle(ctx, a.Slug)
}

func (uc *SocialUsecase) AddComment(ctx context.Context, in *Comment) (rv *Comment, err error) {
	return uc.cr.AddComment(ctx, in)
}

func (uc *SocialUsecase) ListComments(ctx context.Context, slug string) (rv []*Comment, err error) {
	uc.cr.ListComments(ctx, slug)
	return nil, nil
}

func (uc *SocialUsecase) DeleteComment(ctx context.Context, id uint) (err error) {
	a, err := uc.cr.GetComment(ctx, id)
	if err != nil {
		return err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).Username) {
		return errors.New("no permission 401")
	}
	err = uc.cr.DeleteComment(ctx, id)
	return err
}

func (uc *SocialUsecase) FeedArticles(ctx context.Context, opts ...ListOption) (rv []*Article, err error) {
	rv, err = uc.ar.ListArticles(ctx, opts...)
	return rv, err
}

func (uc *SocialUsecase) ListArticles(ctx context.Context, opts ...ListOption) (rv []*Article, err error) {
	rv, err = uc.ar.ListArticles(ctx, opts...)
	return rv, err
}

func (uc *SocialUsecase) UpdateArticle(ctx context.Context, in *Article) (rv *Article, err error) {
	a, err := uc.ar.GetArticle(ctx, in.Slug)
	if err != nil {
		return nil, err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).Username) {
		return nil, errors.New("no permission 401")
	}
	rv, err = uc.ar.UpdateArticle(ctx, in)
	return nil, nil
}

func (uc *SocialUsecase) GetTags(ctx context.Context) (rv []*Tag, err error) {
	uc.tr.GetTags(ctx)
	return nil, nil
}

func (uc *SocialUsecase) FavoriteArticle(ctx context.Context, slug string) (rv *Article, err error) {
	a, err := uc.ar.GetArticle(ctx, slug)
	if err != nil {
		return nil, err
	}
	cu := auth.FromContext(ctx)
	err = uc.ar.FavoriteArticle(ctx, cu.Username, a.Slug)
	if err != nil {
		return nil, err
	}
	return
}

func (uc *SocialUsecase) UnfavoriteArticle(ctx context.Context, slug string) (rv *Article, err error) {
	a, err := uc.ar.GetArticle(ctx, slug)
	if err != nil {
		return nil, err
	}
	cu := auth.FromContext(ctx)
	err = uc.ar.UnfavoriteArticle(ctx, cu.Username, a.Slug)
	if err != nil {
		return nil, err
	}
	return
}
