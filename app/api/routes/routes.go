package routes

import (
	"go.uber.org/fx"
	"new-blog/app/api/types"
	"new-blog/core/http"
)

type Routes struct {
	fx.In
	Http *http.Service
}

func NewRoutes(t Routes, h *http.Service) *types.ApiRouter {
	return &types.ApiRouter{
		RouterGroup: t.Http.Gin.Group("/api"),
	}
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
	// Add injectors here
)
