package rpc

import (
	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/application/in"
	"github.com/okm321/mahking/go/internal/application/out"
	"github.com/okm321/mahking/go/internal/domain"
	groupv1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/group/v1"
)

func toCreateGroupInput(req *groupv1.CreateGroupRequest) in.CreateGroupWithRule {
	rules := req.GetRules()
	input := in.CreateGroupWithRule{
		Name:        req.GetName(),
		MemberNames: req.GetMemberNames(),
		Rules: in.Rules{
			MahjongType:           domain.MahjongType(rules.GetMahjongType()),
			InitialPoints:         int(rules.GetInitialPoints()),
			ReturnPoints:          int(rules.GetReturnPoints()),
			RankingPointsFirst:    int(rules.GetRankingPointsFirst()),
			RankingPointsSecond:   int(rules.GetRankingPointsSecond()),
			RankingPointsThird:    int(rules.GetRankingPointsThird()),
			FractionalCalculation: domain.FractionalCalculation(rules.GetFractionalCalculation()),
			UseBust:               rules.GetUseBust(),
			UseChip:               rules.GetUseChip(),
		},
	}

	if rules != nil {
		input.Rules.RankingPointsFour = nullIntFromPtr(rules.RankingPointsFour)
		input.Rules.BustPoint = nullIntFromPtr(rules.BustPoint)
		input.Rules.ChipPoint = nullIntFromPtr(rules.ChipPoint)
	}
	return input
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
