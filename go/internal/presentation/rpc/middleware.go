package rpc

import (
	"net/http"

	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

func traceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pkgtrace.SpanFromRemote(r.Context(), r.Header)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
