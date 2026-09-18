package rpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/application"
	"github.com/okm321/mahking/go/internal/application/in"
	"github.com/okm321/mahking/go/internal/application/out"
	"github.com/okm321/mahking/go/internal/domain"
	groupv1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1"
	"github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1/groupv1connect"
	pkgtrace "github.com/okm321/mahking/go/pkg/trace"
)

type GroupServer struct {
	groupv1connect.UnimplementedGroupServiceHandler
	usecase *application.GroupUsecase
}

func NewGroupServer(usecase *application.GroupUsecase) *GroupServer {
	return &GroupServer{
		usecase: usecase,
	}
}

func (s *GroupServer) GetGroup(ctx context.Context, req *connect.Request[groupv1.GetGroupRequest]) (_ *connect.Response[groupv1.GetGroupResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupServer.GetGroup")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	group, err := s.usecase.Get(ctx, req.Msg.Uid)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&groupv1.GetGroupResponse{
		Group: toProtoGroup(group),
	}), nil
}

func (s *GroupServer) CreateGroup(ctx context.Context, req *connect.Request[groupv1.CreateGroupRequest]) (_ *connect.Response[groupv1.CreateGroupResponse], err error) {
	ctx = pkgtrace.StartSpan(ctx, "GroupServer.CreateGroup")
	defer func() { pkgtrace.EndSpan(ctx, err) }()

	rules := req.Msg.GetRules()
	group, err := s.usecase.Create(ctx, in.CreateGroupWithRule{
		Name:        req.Msg.GetName(),
		MemberNames: req.Msg.GetMemberNames(),
		Rules: in.Rules{
			MahjongType:           domain.MahjongType(rules.GetMahjongType()),
			InitialPoints:         int(rules.GetInitialPoints()),
			ReturnPoints:          int(rules.GetReturnPoints()),
			RankingPointsFirst:    int(rules.GetRankingPointsFirst()),
			RankingPointsSecond:   int(rules.GetRankingPointsSecond()),
			RankingPointsThird:    int(rules.GetRankingPointsThird()),
			RankingPointsFour:     nullIntFromPtr(rules.RankingPointsFour),
			FractionalCalculation: domain.FractionalCalculation(rules.GetFractionalCalculation()),
			UseBust:               rules.GetUseBust(),
			BustPoint:             nullIntFromPtr(rules.BustPoint),
			UseChip:               rules.GetUseChip(),
			ChipPoint:             nullIntFromPtr(rules.ChipPoint),
		},
	})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&groupv1.CreateGroupResponse{
		Group: toProtoGroup(group),
	}), nil
}

func toProtoGroup(g *out.Group) *groupv1.Group {
	return &groupv1.Group{
		Id:   g.ID,
		Uid:  g.UID,
		Name: g.Name,
	}
}

func nullIntFromPtr(p *int32) null.Int {
	if p == nil {
		return null.Int{}
	}
	return null.IntFrom(int64(*p))
}
