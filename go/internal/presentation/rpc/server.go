package rpc

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/okm321/mahking/go/config"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/game/v1/gamev1connect"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1/groupv1connect"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
	"github.com/okm321/mahking/go/pkg/logger"
)

type ServerSet struct {
	Group *GroupServer
	Game  *GameServer
}

func NewHandler(servers ServerSet) http.Handler {
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(errorInterceptor(), validate.NewInterceptor())

	mux.Handle(groupv1connect.NewGroupServiceHandler(servers.Group, interceptors))
	mux.Handle(gamev1connect.NewGameServiceHandler(servers.Game, interceptors))
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return traceContextMiddleware(accessLogMiddleware(mux))
}

func Run(cfg *config.Config, handler http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port),
		Handler:      handler,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// 別goroutineでサーバー起動
	go func() {
		logger.InfofContext(context.Background(), "starting rpc server on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.InfofContext(context.Background(), "shutting down the server: %v", err)
		}
	}()

	// シグナル待機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	<-quit

	logger.InfoContext(context.Background(), "Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return pkgerror.Errorf("graceful shut down failed: %w", err)
	}

	logger.InfoContext(ctx, "graceful shutdown success")

	return nil
}
