package main

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

func readyz(w http.ResponseWriter, r *http.Request) error {
	code := http.StatusOK
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
	return nil
}

func livez(w http.ResponseWriter, r *http.Request) error {
	code := http.StatusOK
	w.WriteHeader(code)
	w.Write([]byte(http.StatusText(code)))
	return nil
}

func get(c *redis.Client) funcErr {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		q := r.URL.Query()
		key := q.Get("key")
		value := c.Get(r.Context(), key)
		if value.Err() != nil {
			return value.Err()
		}
		res, err := value.Result()
		if err != nil {
			return err
		}
		return encode(w, r, http.StatusOK, res)
	}
	return fn
}
