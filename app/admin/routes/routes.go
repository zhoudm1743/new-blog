package routes

import (
	"go.uber.org/fx"
	"new-blog/app/admin/routes/cms"
	"new-blog/app/admin/routes/user"
	"new-blog/app/admin/types"
	"new-blog/core/http"
)

type Routes struct {
	fx.In
	Http *http.Service
}

func NewRoutes(t Routes, h *http.Service) *types.AdminRouter {
	return &types.AdminRouter{
		RouterGroup: t.Http.Gin.Group("/admin"),
	}
}

var Module = fx.Options(
	fx.Provide(NewRoutes),

	fx.Invoke(cms.ArticleRoutes),
	fx.Invoke(cms.CategoryRoutes),
	fx.Invoke(cms.TagRoutes),
	fx.Invoke(cms.LinkRoutes),
	fx.Invoke(user.UserRoutes), // Add injectors here
)
