package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okm321/mahking/go/config"
	"github.com/okm321/mahking/go/internal/application"
	"github.com/okm321/mahking/go/internal/infrastructure/postgres"
	"github.com/okm321/mahking/go/internal/presentation/rpc"
	pkgpostgres "github.com/okm321/mahking/go/pkg/postgres"
)

type App struct {
	cfg     *config.Config
	handler http.Handler
	pool    *pgxpool.Pool
}

// NewApp creates all dependencies and returns a runnable App.
func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	// Infra
	pool, err := pkgpostgres.Connect(cfg.DBPostgres)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	// Repository
	groupRepo := postgres.NewGroupRepository(pool)

	// Usecase
	groupUsecase := application.NewGroupUsecase(&application.NewGroupUsecaseArgs{
		GroupRepo: groupRepo,
	})

	// Handler
	handler := rpc.NewHandler(rpc.ServerSet{
		Group: rpc.NewGroupServer(groupUsecase),
	})

	return &App{cfg: cfg, handler: handler, pool: pool}, nil
}

// Run starts the RPC server. It blocks until shutdown.
func (a *App) Run() error {
	return rpc.Run(a.cfg, a.handler)
}

// Close releases resources.
func (a *App) Close() {
	if a.pool != nil {
		a.pool.Close()
	}
}
