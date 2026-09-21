package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/guregu/null/v6"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okm321/mahking/go/internal/domain"
	"github.com/okm321/mahking/go/internal/infrastructure/postgres/sqlc"
	pkgerror "github.com/okm321/mahking/go/pkg/error"
	pkgpostgres "github.com/okm321/mahking/go/pkg/postgres"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GroupRepository struct {
	pool *pgxpool.Pool
}

func NewGroupRepository(pool *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{
		pool: pool,
	}
}

func (r *GroupRepository) List(ctx context.Context) (_ []domain.Group, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupRepository.List")
	defer func() { pkgtrace.EndSpan(ctx, err) }()
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))
	rows, err := q.ListGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("list groups: %w", err)
	}

	groups := make([]domain.Group, 0, len(rows))
	for _, row := range rows {
		groups = append(groups, domain.Group{
			ID:   row.ID,
			UID:  row.Uid,
			Name: row.Name,
		})
	}
	return groups, nil
}

func (r *GroupRepository) GetByUUID(ctx context.Context, uid string) (_ *domain.Group, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupRepository.GetByUUID")
	defer func() { pkgtrace.EndSpan(ctx, err) }()
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))
	row, err := q.GetGroupByID(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pkgerror.WithStack(pkgerror.ErrNotFound)
		}
		return nil, pkgerror.Wrap(err, "get group by uid")
	}

	members, err := q.ListMembersByGroupID(ctx, row.ID)
	if err != nil {
		return nil, pkgerror.Wrap(err, "list members by group id")
	}

	rule, err := q.GetRuleByGroupID(ctx, row.ID)
	if err != nil {
		return nil, pkgerror.Wrap(err, "get rule by group id")
	}

	dms := make([]*domain.Member, 0, len(members))
	for _, m := range members {
		dms = append(dms, &domain.Member{
			ID:      m.ID,
			GroupID: m.GroupID,
			Name:    m.Name,
		})
	}

	return &domain.Group{
		ID:      row.ID,
		UID:     row.Uid,
		Name:    row.Name,
		Members: dms,
		Rule:    toDomainRule(rule),
	}, nil
}

func (r *GroupRepository) Create(ctx context.Context, group *domain.Group) (_ *domain.Group, err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupRepository.Create")
	defer func() { pkgtrace.EndSpan(ctx, err) }()
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))
	row, err := q.CreateGroup(ctx, group.Name)
	if err != nil {
		return nil, err
	}

	group.ID = row.ID

	err = r.createRelatedInfo(ctx, group)
	if err != nil {
		return nil, err
	}

	return &domain.Group{
		ID:   row.ID,
		UID:  row.Uid,
		Name: row.Name,
	}, nil
}

func (r *GroupRepository) createRelatedInfo(ctx context.Context, group *domain.Group) (err error) {
	q := sqlc.New(pkgpostgres.GetExecutor(ctx, r.pool))
	memberParams := make([]sqlc.CreateMembersParams, 0, len(group.Members))
	for _, m := range group.Members {
		memberParams = append(memberParams, sqlc.CreateMembersParams{
			GroupID: group.ID,
			Name:    m.Name,
		})
	}
	_, err = q.CreateMembers(ctx, memberParams)
	if err != nil {
		return pkgerror.Wrap(err, "create related members")
	}

	ruleParam := sqlc.CreateRuleParams{
		GroupID:               group.ID,
		MahjongType:           int32(group.Rule.MahjongType),         //nolint:gosec // 麻雀タイプは1-2の範囲
		InitialPoints:         int32(group.Rule.InitialPoints),       //nolint:gosec // 点数はint32範囲内
		ReturnPoints:          int32(group.Rule.ReturnPoints),        //nolint:gosec // 点数はint32範囲内
		RankingPointsFirst:    int32(group.Rule.RankingPointsFirst),  //nolint:gosec // 点数はint32範囲内
		RankingPointsSecond:   int32(group.Rule.RankingPointsSecond), //nolint:gosec // 点数はint32範囲内
		RankingPointsThird:    int32(group.Rule.RankingPointsThird),  //nolint:gosec // 点数はint32範囲内
		RankingPointsFourth:   null.IntFromPtr(group.Rule.RankingPointsFour.Ptr()),
		FractionalCalculation: int32(group.Rule.FractionalCalculation), //nolint:gosec // 計算方法は1-5の範囲
		UseBust:               group.Rule.UseBust,
		BustPoint:             null.IntFromPtr(group.Rule.BustPoint.Ptr()),
		UseChip:               group.Rule.UseChip,
		ChipPoint:             null.IntFromPtr(group.Rule.ChipPoint.Ptr()),
	}
	_, err = q.CreateRule(ctx, ruleParam)
	if err != nil {
		return pkgerror.Wrap(err, "create related rule")
	}

	return nil
}

var _ domain.GroupRepository = (*GroupRepository)(nil)

func toDomainRule(r sqlc.Rule) *domain.Rule {
	return &domain.Rule{
		ID:                    r.ID,
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
	}
}
