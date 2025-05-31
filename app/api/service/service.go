package service

import (
	"go.uber.org/fx"
	"new-blog/app/api/service/user"
)

var Module = fx.Options(
	// Provide your dependencies here
	fx.Provide(user.NewUserService),
)
