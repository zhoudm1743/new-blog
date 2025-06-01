package service

import (
    "go.uber.org/fx"
)

var Module = fx.Options(
    // Provide your dependencies here
	fx.Provide(user.NewUserService),
)