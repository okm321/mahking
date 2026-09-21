package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okm321/mahking/go/internal/domain"
	"github.com/okm321/mahking/go/internal/infrastructure/postgres/sqlc"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
	pkgpostgres "github.com/okm321/mahking/go/pkg/postgres"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GameRepository struct {
	pool *pgxpool.Pool
}

func NewGameRepository(pool *pgxpool.Pool) *GameRepository {
	return &GameRepository{
		pool: pool,
	}
}

func (r *GameRepository) Create(ctx context.Context, game *domain.Game) (_ *domain.Game, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameRepository.Create")
	defer func() { pkgtrace.EndSpan(ctx, err) }()
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))

	row, err := q.CreateGame(ctx, sqlc.CreateGameParams{
		GroupID:  game.GroupID,
		Note:     pgtype.Text{String: game.Note, Valid: game.Note != ""},
		PlayedAt: pgtype.Timestamptz{Time: game.PlayedAt, Valid: true},
	})
	if err != nil {
		return nil, pkgerror.Wrap(err, "create game")
	}

	game.ID = row.ID
	game.GameRule.GameID = row.ID
	for _, s := range game.GameScores {
		s.GameID = row.ID
		s.GroupID = game.GroupID
	}

	if err := r.createRelatedInfo(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (r *GameRepository) createRelatedInfo(ctx context.Context, game *domain.Game) error {
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))

	gr := game.GameRule
	ruleRow, err := q.CreateGameRule(ctx, sqlc.CreateGameRuleParams{
		GameID:                game.ID,
		GroupID:               game.GroupID,
		MahjongType:           int32(gr.MahjongType),         //nolint:gosec // 麻雀タイプは1-2の範囲
		InitialPoints:         int32(gr.InitialPoints),       //nolint:gosec // 点数はint32範囲内
		ReturnPoints:          int32(gr.ReturnPoints),        //nolint:gosec // 点数はint32範囲内
		RankingPointsFirst:    int32(gr.RankingPointsFirst),  //nolint:gosec // 点数はint32範囲内
		RankingPointsSecond:   int32(gr.RankingPointsSecond), //nolint:gosec // 点数はint32範囲内
		RankingPointsThird:    int32(gr.RankingPointsThird),  //nolint:gosec // 点数はint32範囲内
		RankingPointsFourth:   gr.RankingPointsFour,
		FractionalCalculation: int32(gr.FractionalCalculation), //nolint:gosec // 計算方法は1-5の範囲
		UseBust:               gr.UseBust,
		BustPoint:             gr.BustPoint,
		UseChip:               gr.UseChip,
		ChipPoint:             gr.ChipPoint,
	})
	if err != nil {
		return pkgerror.Wrap(err, "create game rule")
	}
	gr.ID = ruleRow.ID

	scoreParams := make([]sqlc.CreateGameScoresParams, 0, len(game.GameScores))
	for _, s := range game.GameScores {
		scoreParams = append(scoreParams, sqlc.CreateGameScoresParams{
			GameID:    game.ID,
			GroupID:   game.GroupID,
			MemberID:  s.MemberID,
			Seat:      int32(s.Seat),     //nolint:gosec // 席は1-4の範囲
			Ranking:   int32(s.Ranking),  //nolint:gosec // 順位は1-4の範囲
			RawScore:  int32(s.RawScore), //nolint:gosec // 素点はint32範囲内
			Point:     s.Point,
			ChipCount: s.ChipCount,
			IsBusted:  s.IsBusted,
		})
	}
	if _, err := q.CreateGameScores(ctx, scoreParams); err != nil {
		return pkgerror.Wrap(err, "create game scores")
	}

	return nil
}

// List グループの対局一覧を新しい順に返す
// ponytail: game ごとに rule と scores を引く N+1。件数が増えたら ANY(@game_ids) でまとめて引く
func (r *GameRepository) List(ctx context.Context, groupID int64) (_ []domain.Game, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GameRepository.List")
	defer func() { pkgtrace.EndSpan(ctx, err) }()
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))

	rows, err := q.ListGamesByGroupID(ctx, groupID)
	if err != nil {
		return nil, pkgerror.Wrap(err, "list games by group id")
	}

	games := make([]domain.Game, 0, len(rows))
	for _, row := range rows {
		ruleRow, err := q.GetGameRuleByGameID(ctx, row.ID)
		if err != nil {
			return nil, pkgerror.Wrapf(err, "get game rule by game id %d", row.ID)
		}
		scoreRows, err := q.ListGameScoresByGameID(ctx, row.ID)
		if err != nil {
			return nil, pkgerror.Wrapf(err, "list game scores by game id %d", row.ID)
		}

		scores := make([]*domain.GameScore, 0, len(scoreRows))
		for _, s := range scoreRows {
			scores = append(scores, &domain.GameScore{
				ID:        s.ID,
				GameID:    s.GameID,
				GroupID:   s.GroupID,
				MemberID:  s.MemberID,
				Seat:      domain.Seat(s.Seat),
				Ranking:   int(s.Ranking),
				RawScore:  int(s.RawScore),
				Point:     s.Point,
				ChipCount: s.ChipCount,
				IsBusted:  s.IsBusted,
			})
		}

		games = append(games, domain.Game{
			ID:         row.ID,
			GroupID:    row.GroupID,
			Note:       row.Note.String,
			PlayedAt:   row.PlayedAt.Time,
			GameRule:   toDomainGameRule(ruleRow),
			GameScores: scores,
		})
	}

	return games, nil
}

func toDomainGameRule(r sqlc.GameRule) *domain.GameRule {
	return &domain.GameRule{
		ID:     r.ID,
		GameID: r.GameID,
		Rule: domain.Rule{
			GroupID:               r.GroupID,
			MahjongType:           domain.MahjongType(r.MahjongType),
			InitialPoints:         int(r.InitialPoints),
			ReturnPoints:          int(r.ReturnPoints),
			RankingPointsFirst:    int(r.RankingPointsFirst),
			RankingPointsSecond:   int(r.RankingPointsSecond),
			RankingPointsThird:    int(r.RankingPointsThird),
			RankingPointsFour:     r.RankingPointsFourth,
			FractionalCalculation: domain.FractionalCalculation(r.FractionalCalculation),
			UseBust:               r.UseBust,
			BustPoint:             r.BustPoint,
			UseChip:               r.UseChip,
			ChipPoint:             r.ChipPoint,
		},
	}
}

var _ domain.GameRepository = (*GameRepository)(nil)
