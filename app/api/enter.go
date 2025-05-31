package api

import (
	"go.uber.org/fx"
	"new-blog/app/api/routes"
	"new-blog/app/api/service"
)

var Module = fx.Module("api",
	// Add your dependencies here
	service.Module,
	routes.Module,
)
