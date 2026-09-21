package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/okm321/mahking/go/internal/application"
	gamev1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/game/v1"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/game/v1/gamev1connect"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GameServer struct {
	gamev1connect.UnimplementedGameServiceHandler
	usecase *application.GameUsecase
}

func NewGameServer(usecase *application.GameUsecase) *GameServer {
	return &GameServer{
		usecase: usecase,
	}
}

func (s *GameServer) CreateGame(ctx context.Context, req *connect.Request[gamev1.CreateGameRequest]) (_ *connect.Response[gamev1.CreateGameResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameServer.CreateGame")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	game, err := s.usecase.Create(ctx, toCreateGameInput(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&gamev1.CreateGameResponse{
		Game: toProtoGame(game),
	}), nil
}

func (s *GameServer) ListGames(ctx context.Context, req *connect.Request[gamev1.ListGamesRequest]) (_ *connect.Response[gamev1.ListGamesResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameServer.ListGames")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	games, err := s.usecase.List(ctx, req.Msg.GetGroupUid())
	if err != nil {
		return nil, err
	}

	res := make([]*gamev1.Game, 0, len(games))
	for i := range games {
		res = append(res, toProtoGame(&games[i]))
	}

	return connect.NewResponse(&gamev1.ListGamesResponse{
		Games: res,
	}), nil
}
