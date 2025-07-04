package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) error {
	url := getenv("REDIS_URL")
	c, err := newClient(ctx, url)
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	addRoutes(mux, c)
	host := getenv("APP_HOST")
	port := getenv("APP_PORT")
	addr := net.JoinHostPort(host, port)
	slog.Info("SUCESSFULLY STARTED SERVER!", "addr", addr)
	return http.ListenAndServe(addr, mux)
}

func newClient(ctx context.Context, url string) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	c := redis.NewClient(opts)
	examples := map[string]string{
		"teccle": "group",
		"etomer": "remote",
	}
	for key, value := range examples {
		err := c.Set(ctx, key, value, 0).Err()
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
