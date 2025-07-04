package main

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

func addRoutes(mux *http.ServeMux, c *redis.Client) {
	mux.Handle("GET /livez", toHandler(livez))
	mux.Handle("GET /readyz", toHandler(readyz))
	mux.Handle("GET /", toHandler(get(c)))
}
