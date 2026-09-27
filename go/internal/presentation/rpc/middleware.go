package rpc

import (
	"net/http"
	"time"

	connectcors "connectrpc.com/cors"
	"github.com/go-chi/cors"

	"github.com/okm321/mahking/go/pkg/logger"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

func traceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pkgtrace.SpanFromRemote(r.Context(), r.Header)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// accessLogMiddleware 1リクエストにつき1行のアクセスログを出す
func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}

		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		start := time.Now()

		next.ServeHTTP(rec, r)

		logger.InfoContext(r.Context(), "http request",
			logger.HTTPAttr(r, rec.status, time.Since(start), rec.size),
		)
	})
}

// statusRecorder レスポンスのステータスとサイズを記録するResponseWriter
type statusRecorder struct {
	http.ResponseWriter
	status int
	size   int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *statusRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.size += n
	return n, err
}

// corsMiddleware ブラウザから直接RPCを叩けるようにCORSヘッダーを付ける。
// 許可するヘッダーとメソッドは connectrpc.com/cors が Connect / gRPC-Web に必要なものを返す。
func corsMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         86400, // 24時間。プリフライトの結果をブラウザがキャッシュする
	})
}
