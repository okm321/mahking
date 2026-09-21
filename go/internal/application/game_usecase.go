package application

import (
	"context"

	appin "github.com/okm321/mahking/go/internal/application/in"
	appout "github.com/okm321/mahking/go/internal/application/out"
	"github.com/okm321/mahking/go/internal/domain"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GameUsecase struct {
	gameRepo  domain.GameRepository
	groupRepo domain.GroupRepository
	tx        domain.Transactioner
}

type NewGameUsecaseArgs struct {
	GameRepo  domain.GameRepository
	GroupRepo domain.GroupRepository
	Tx        domain.Transactioner
}

func NewGameUsecase(args *NewGameUsecaseArgs) *GameUsecase {
	return &GameUsecase{
		gameRepo:  args.GameRepo,
		groupRepo: args.GroupRepo,
		tx:        args.Tx,
	}
}

func (u *GameUsecase) Create(ctx context.Context, in appin.CreateGame) (_ *appout.Game, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameUsecase.Create")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	group, err := u.groupRepo.GetByUUID(ctx, in.GroupUID)
	if err != nil {
		return nil, err
	}

	memberIDs := make(map[int64]struct{}, len(group.Members))
	for _, m := range group.Members {
		memberIDs[m.ID] = struct{}{}
	}

	scores := make([]*domain.GameScore, 0, len(in.Scores))
	for _, s := range in.Scores {
		if _, ok := memberIDs[s.MemberID]; !ok {
			return nil, pkgerror.NewClientErrorf("グループに存在しないメンバーです: メンバーID %d", s.MemberID)
		}
		scores = append(scores, &domain.GameScore{
			GroupID:   group.ID,
			MemberID:  s.MemberID,
			Seat:      s.Seat,
			Ranking:   s.Ranking,
			RawScore:  s.RawScore,
			ChipCount: s.ChipCount,
			IsBusted:  s.IsBusted,
		})
	}

	game, err := domain.NewGame(ctx, group.ID, domain.NewGameArgs{
		Note:       in.Note,
		PlayedAt:   in.PlayedAt,
		GameRule:   domain.NewGameRuleFromRule(group.Rule),
		GameScores: scores,
	})
	if err != nil {
		return nil, err
	}

	var created *domain.Game
	err = domain.WithTransaction(ctx, u.tx, func(ctx context.Context) error {
		var createErr error
		created, createErr = u.gameRepo.Create(ctx, game)
		return createErr
	})
	if err != nil {
		return nil, err
	}

	out := appout.NewGame(*created)
	return &out, nil
}

func (u *GameUsecase) List(ctx context.Context, groupUID string) (_ []appout.Game, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameUsecase.List")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	group, err := u.groupRepo.GetByUUID(ctx, groupUID)
	if err != nil {
		return nil, err
	}

	games, err := u.gameRepo.List(ctx, group.ID)
	if err != nil {
		return nil, err
	}

	res := make([]appout.Game, 0, len(games))
	for _, g := range games {
		res = append(res, appout.NewGame(g))
	}
	return res, nil
}
