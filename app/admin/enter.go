package admin

import (
	"go.uber.org/fx"
	"new-blog/app/admin/routes"
	"new-blog/app/admin/service"
)

var Module = fx.Module("admin",
	// Add your dependencies here
	service.Module,
	routes.Module,
)
