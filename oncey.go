// Package oncey provides an HTTP middleware for idempotency keys.
// It ensures requests with the same Idempotency-Key are processed only once.
package oncey

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sokchamrernheng/oncey/internal/capture"
	"github.com/sokchamrernheng/oncey/internal/store"
)

type Option struct {
	TTL   time.Duration
	Store store.Store
}

func NewHTTPMiddleware(opt Option) func(http.Handler) http.Handler {
	ttl := opt.TTL
	st := opt.Store

	if ttl == 0 {
		ttl = 1800 * time.Second
	}
	if st == nil {
		st = store.NewMemoryStore()
	}
	return func(next http.Handler) http.Handler {
		return idempotencyMiddleware(next, ttl, st)
	}
}

func idempotencyMiddleware(next http.Handler, ttl time.Duration, st store.Store) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		key, err := extractKey(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cw := capture.NewHttpCaptureWriter(w)
		callback := func(ctx context.Context) error {
			next.ServeHTTP(cw, r)
			return nil 
		}

		DoOnce(
			r.Context(), 
			key, 
			ttl,
			st,
			callback,
		)
	}
	return http.HandlerFunc(fn)
}

func extractKey(r *http.Request) (string, error) {
	key := r.Header.Get("X-Idempotency-Key")

	if key == "" {
		return "", errors.New("missing idempotency key")
	}

	return key, nil
}

type Callback func(ctx context.Context) error

func DoOnce(ctx context.Context, key string,ttl time.Duration, store store.Store, fn Callback) error {
	
	// check store
	// cachedResp, err := store.Get(key)
	// if (err != nil) {
	// 	//process error
	// }
	// if already done → return
	// mark in-progress


	// run fn
	// on success → mark done
	fn(ctx)
	return nil
}
