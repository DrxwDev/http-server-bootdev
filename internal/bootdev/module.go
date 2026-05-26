// Package bootdev
package bootdev

import "go.uber.org/fx"

var Module = fx.Module(
	"bootdev",
	fx.Provide(
		NewBootController,
		NewBootService,
	),
)
