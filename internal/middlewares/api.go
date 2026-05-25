package middlewares

import "sync/atomic"

type APIConfig struct {
	fileserverHits atomic.Int32
}

func NewAPIConfig() *APIConfig {
	return &APIConfig{
		fileserverHits: atomic.Int32{},
	}
}
