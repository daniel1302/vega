package http

import (
	"net/http"
	"strings"

	"github.com/rs/cors"
)

// CORSConfig represents the configuration for CORS.
type CORSConfig struct {
	AllowedOrigins string `long:"allowed-origins" description:"Allowed origins for CORS"`
	MaxAge         int    `long:"max-age" description:"Max age (in seconds) for preflight cache"`
}

func CORSOptions(config CORSConfig) cors.Options {
	return cors.Options{
		AllowOriginFunc: AllowedOrigin(config.AllowedOrigins),
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		MaxAge:           config.MaxAge,
		AllowCredentials: false,
	}
}

func AllowedOrigin(allowedOrigins string) func(origin string) bool {
	return func(origin string) bool {
		origin = strings.TrimPrefix(strings.TrimPrefix(origin, "https://"), "http://")
		return allowedOrigins == "" || allowedOrigins == "*" || strings.Contains(allowedOrigins, origin)
	}
}
