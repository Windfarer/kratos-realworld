package service

import (
	"context"

	v1 "kratos-realworld/api/conduit/v1"
	"kratos-realworld/internal/biz"
)

func convertSingleArticleReply(do *biz.Article) *v1.SingleArticleReply {
	return &v1.SingleArticleReply{
		Article: &v1.SingleArticleReply_Article{
			Slug:           do.Slug,
			Title:          do.Title,
			Description:    do.Description,
			Body:           do.Body,
			TagList:        do.TagList,
			CreatedAt:      do.CreatedAt.String(),
			UpdatedAt:      do.UpdatedAt.String(),
			Favorited:      do.Favorited,
			FavoritesCount: do.FavoritesCount,
			Author: &v1.SingleArticleReply_Author{
				Username:  do.Author.Username,
				Bio:       do.Author.Bio,
				Image:     do.Author.Image,
				Following: do.Author.Following,
			},
		},
	}
}

func convertMultipleArticlesReply(dos []*biz.Article) (dto *v1.MultipleArticlesReply) {
	dto = &v1.MultipleArticlesReply{
		Articles: make([]*v1.MultipleArticlesReply_Articles, 0),
	}
	for _, do := range dos {
		dto.Articles = append(dto.Articles, &v1.MultipleArticlesReply_Articles{
			Slug:           do.Slug,
			Title:          do.Title,
			Description:    do.Description,
			Body:           do.Body,
			TagList:        do.TagList,
			CreatedAt:      do.CreatedAt.String(),
			UpdatedAt:      do.UpdatedAt.String(),
			Favorited:      do.Favorited,
			FavoritesCount: do.FavoritesCount,
			Author: &v1.MultipleArticlesReply_Author{
				Username:  do.Author.Username,
				Bio:       do.Author.Bio,
				Image:     do.Author.Image,
				Following: do.Author.Following,
			},
		},
		)
	}
	return dto
}

func (s *ConduitService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.GetProfile(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{
		Profile: &v1.ProfileReply_Profile{
			Username:  rv.Username,
			Bio:       rv.Bio,
			Image:     rv.Image,
			Following: rv.Following,
		},
	}, nil
}

func (s *ConduitService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.FollowUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{
		Profile: &v1.ProfileReply_Profile{
			Username:  rv.Username,
			Bio:       rv.Bio,
			Image:     rv.Image,
			Following: rv.Following,
		},
	}, nil
}

func (s *ConduitService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.UnfollowUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{
		Profile: &v1.ProfileReply_Profile{
			Username:  rv.Username,
			Bio:       rv.Bio,
			Image:     rv.Image,
			Following: rv.Following,
		},
	}, nil
}

func (s *ConduitService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (reply *v1.SingleArticleReply, err error) {
	rv, err := s.sc.GetArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return convertSingleArticleReply(rv), nil
}

func (s *ConduitService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	rv, err := s.sc.CreateArticle(ctx, &biz.Article{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		TagList:     req.Article.TagList,
	})
	if err != nil {
		return nil, err
	}

	return convertSingleArticleReply(rv), nil
}

func (s *ConduitService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (reply *v1.SingleArticleReply, err error) {
	rv, err := s.sc.UpdateArticle(ctx, &biz.Article{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		TagList:     req.Article.TagList,
	})
	if err != nil {
		return nil, err
	}

	return convertSingleArticleReply(rv), nil
}

func (s *ConduitService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	err = s.sc.DeleteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return convertSingleArticleReply(&biz.Article{Slug: req.Slug}), nil
}

func (s *ConduitService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	reply = &v1.SingleCommentReply{}

	return reply, nil
}

func (s *ConduitService) GetComments(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.MultipleCommentsReply, err error) {
	reply = &v1.MultipleCommentsReply{}

	return reply, nil
}

func (s *ConduitService) DeleteComment(ctx context.Context, req *v1.DeleteCommentRequest) (reply *v1.SingleCommentReply, err error) {
	reply = &v1.SingleCommentReply{}

	return reply, nil
}

func (s *ConduitService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	rv, err := s.sc.ListArticles(ctx)
	if err != nil {
		return nil, err
	}

	return convertMultipleArticlesReply(rv), nil
}

func (s *ConduitService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	rv, err := s.sc.ListArticles(ctx)
	if err != nil {
		return nil, err
	}

	return convertMultipleArticlesReply(rv), nil
}

func (s *ConduitService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	reply = &v1.TagListReply{}

	return reply, nil
}

func (s *ConduitService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	rv, err := s.sc.FavoriteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	return convertSingleArticleReply(rv), nil
}

func (s *ConduitService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	rv, err := s.sc.UnfavoriteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	return convertSingleArticleReply(rv), nil
}
