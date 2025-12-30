package oncey

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOnceyBasic(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	option := Option{

	}
	mw := NewHTTPMiddleware(option)(handler)

	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("X-Idempotency-Key", "33cfcdae-1d28-47a1-8075-f0211867e87e")

	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	if rec.Body.String() != "ok" {
		t.Fatalf("expected ok, got %q", rec.Body.String())
	}
}