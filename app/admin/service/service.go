package service

import (
	"go.uber.org/fx"
	"new-blog/app/admin/service/cms"
	"new-blog/app/admin/service/user"
)

var Module = fx.Options(
	// Provide your dependencies here
	fx.Provide(user.NewUserService),
	fx.Provide(cms.NewLinkService),
	fx.Provide(cms.NewTagService),
	fx.Provide(cms.NewCategoryService),
	fx.Provide(cms.NewArticleService),
)
