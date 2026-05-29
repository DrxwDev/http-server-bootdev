// Package chirpy
package chirpy

import "go.uber.org/fx"

var Module = fx.Module(
	"chirpy",
	fx.Provide(
		NewChirpRepository,
		NewChirpService,
		NewChirpController,
	),
	fx.Invoke(
		ChirpRoutes,
	),
)
