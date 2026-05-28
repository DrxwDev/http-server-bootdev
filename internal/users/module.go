// Package users
package users

import "go.uber.org/fx"

var Module = fx.Module(
	"users",
	fx.Provide(
		NewUserRepository,
		NewUserService,
		NewUserController,
	),
	fx.Invoke(
		UserRoutes,
	),
)
