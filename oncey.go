// Package oncey provides an HTTP middleware for idempotency keys.
// It ensures requests with the same Idempotency-Key are processed only once.
package oncey

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sokchamrernheng/oncey/internal/capture"
	err "github.com/sokchamrernheng/oncey/internal/errors"
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
		key, error := extractKey(r)

		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}

		value, error := st.Get(key)

		if error == nil {
			w.Write(value)
		}
		if errors.Is(error, err.ErrNotFound) {
			cw := capture.NewHttpCaptureWriter(w)
			next.ServeHTTP(cw, r)

			result := cw.GetResult()
			st.Set(key, result, ttl)

			fmt.Println(st.Get(key))
		} else {
			//real error
			//maybe some error or panic
		}
	}
	return http.HandlerFunc(fn)
}

func extractKey(r *http.Request) (string, error) {
	key := r.Header.Get("X-Idempotency-Key")

	if key == "" {
		return "", err.ErrKeyNotSet
	}

	return key, nil
}
