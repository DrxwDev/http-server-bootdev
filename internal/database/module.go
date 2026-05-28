// Package database
package database

import (
	"log"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",
	fx.Provide(
		NewPool,
		NewQueries,
	),
	fx.Invoke(
		func() {
			log.Println("Connected to postgres")
		},
	),
)
