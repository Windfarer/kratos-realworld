package biz

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-realworld/internal/pkg/middleware/auth"
)

type ArticleRepo interface {
	List(ctx context.Context, opts ...ListOption) ([]*Article, error)
	Get(ctx context.Context, slug string) (*Article, error)
	Create(ctx context.Context, a *Article) (*Article, error)
	Update(ctx context.Context, a *Article) (*Article, error)
	Delete(ctx context.Context, a *Article) error
	GetArticle(ctx context.Context, aid uint) (*Article, error)

	Favorite(ctx context.Context, currentUserID uint, aid uint) error
	Unfavorite(ctx context.Context, currentUserID uint, aid uint) error
	GetFavoritesStatus(ctx context.Context, currentUserID uint, as []*Article) (favorited []bool, err error)

	ListTags(ctx context.Context) ([]Tag, error)
}

type CommentRepo interface {
	Create(ctx context.Context, c *Comment) (*Comment, error)
	Get(ctx context.Context, id uint) (*Comment, error)
	List(ctx context.Context, slug string) ([]*Comment, error)
	Delete(ctx context.Context, id uint) error
}

type SocialUsecase struct {
	ar ArticleRepo
	cr CommentRepo
	pr ProfileRepo

	log *log.Helper
}

type Article struct {
	ID             uint
	Slug           string
	Title          string
	Description    string
	Body           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	TagList        []string
	Favorited      bool
	FavoritesCount uint32

	AuthorUserID uint

	Author *Profile
}

type Comment struct {
	ID        uint
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time

	ArticleID uint

	Article  *Article
	AuthorID uint
	Author   *Profile
}

type Tag string

func slugify(title string) string {
	re, _ := regexp.Compile(`[^\w]`)
	return strings.ToLower(re.ReplaceAllString(title, "-"))
}

func (o *Article) verifyAuthor(id uint) bool {
	return o.Author.ID == id
}

func (o *Comment) verifyAuthor(id uint) bool {
	return o.Author.ID == id
}

func NewSocialUsecase(
	ar ArticleRepo,
	pr ProfileRepo,
	cr CommentRepo,
	logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, pr: pr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) GetProfile(ctx context.Context, username string) (rv *Profile, err error) {
	return uc.pr.GetProfile(ctx, username)
}

func (uc *SocialUsecase) FollowUser(ctx context.Context, username string) (rv *Profile, err error) {
	cu := auth.FromContext(ctx)

	fu, err := uc.pr.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	err = uc.pr.FollowUser(ctx, cu.UserID, fu.ID)
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
	fu, err := uc.pr.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	uc.pr.UnfollowUser(ctx, cu.UserID, fu.ID)
	if err != nil {
		return nil, err
	}
	rv, err = uc.pr.GetProfile(ctx, username)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *SocialUsecase) GetArticle(ctx context.Context, slug string) (rv *Article, err error) {
	return uc.ar.Get(ctx, slug)
}

func (uc *SocialUsecase) CreateArticle(ctx context.Context, in *Article) (rv *Article, err error) {
	u := auth.FromContext(ctx)
	in.Slug = slugify(in.Title)
	in.AuthorUserID = u.UserID
	a, err := uc.ar.Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return a, err
}

func (uc *SocialUsecase) DeleteArticle(ctx context.Context, slug string) (err error) {
	a, err := uc.ar.Get(ctx, slug)
	if err != nil {
		return err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).UserID) {
		return errors.New("no permission 401")
	}
	return uc.ar.Delete(ctx, a)
}

func (uc *SocialUsecase) AddComment(ctx context.Context, slug string, in *Comment) (rv *Comment, err error) {
	u := auth.FromContext(ctx)
	in.AuthorID = u.UserID
	in.Article = &Article{Slug: slug}
	return uc.cr.Create(ctx, in)
}

func (uc *SocialUsecase) ListComments(ctx context.Context, slug string) (rv []*Comment, err error) {
	return uc.cr.List(ctx, slug)
}

func (uc *SocialUsecase) DeleteComment(ctx context.Context, id uint) (err error) {
	a, err := uc.cr.Get(ctx, id)
	if err != nil {
		return err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).UserID) {
		return errors.New("no permission 401")
	}
	err = uc.cr.Delete(ctx, id)
	return err
}

func (uc *SocialUsecase) FeedArticles(ctx context.Context, opts ...ListOption) (rv []*Article, err error) {
	rv, err = uc.ar.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *SocialUsecase) ListArticles(ctx context.Context, opts ...ListOption) (rv []*Article, err error) {
	rv, err = uc.ar.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *SocialUsecase) UpdateArticle(ctx context.Context, in *Article) (rv *Article, err error) {
	a, err := uc.ar.Get(ctx, in.Slug)
	if err != nil {
		return nil, err
	}
	if !a.verifyAuthor(auth.FromContext(ctx).UserID) {
		return nil, errors.New("no permission 401")
	}
	rv, err = uc.ar.Update(ctx, in)
	return nil, nil
}

func (uc *SocialUsecase) GetTags(ctx context.Context) (rv []Tag, err error) {
	return uc.ar.ListTags(ctx)
}

func (uc *SocialUsecase) FavoriteArticle(ctx context.Context, slug string) (rv *Article, err error) {
	a, err := uc.ar.Get(ctx, slug)
	if err != nil {
		return nil, err
	}
	cu := auth.FromContext(ctx)
	err = uc.ar.Favorite(ctx, cu.UserID, a.ID)
	if err != nil {
		return nil, err
	}
	a, err = uc.ar.GetArticle(ctx, a.ID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (uc *SocialUsecase) UnfavoriteArticle(ctx context.Context, slug string) (rv *Article, err error) {
	a, err := uc.ar.Get(ctx, slug)
	if err != nil {
		return nil, err
	}
	cu := auth.FromContext(ctx)
	err = uc.ar.Unfavorite(ctx, cu.UserID, a.ID)
	if err != nil {
		return nil, err
	}
	return a, nil
}
