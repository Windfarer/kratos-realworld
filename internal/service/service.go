package service

import (
	"github.com/google/wire"

	v1 "kratos-realworld/api/conduit/v1"
	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewConduitService)

type ConduitService struct {
	v1.UnimplementedConduitServer

	uc  *biz.UserUsecase
	sc  *biz.SocialUsecase
	log *log.Helper
}

func NewConduitService(uc *biz.UserUsecase, sc *biz.SocialUsecase, logger log.Logger) *ConduitService {
	return &ConduitService{uc: uc, sc:sc, log: log.NewHelper(logger)}
}
