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
	members := make([]*groupv1.Member, 0, len(g.Members))
	for _, m := range g.Members {
		members = append(members, &groupv1.Member{Id: m.ID, Name: m.Name})
	}
	return &groupv1.Group{
		Id:      g.ID,
		Uid:     g.UID,
		Name:    g.Name,
		Members: members,
		Rules:   toProtoRules(g.Rule),
	}
}

func toProtoRules(r out.Rule) *groupv1.Rules {
	return &groupv1.Rules{
		MahjongType:           groupv1.MahjongType(r.MahjongType), //nolint:gosec // 麻雀タイプは1-2の範囲
		InitialPoints:         int32(r.InitialPoints),             //nolint:gosec // 点数はint32範囲内
		ReturnPoints:          int32(r.ReturnPoints),              //nolint:gosec // 点数はint32範囲内
		RankingPointsFirst:    int32(r.RankingPointsFirst),        //nolint:gosec // 点数はint32範囲内
		RankingPointsSecond:   int32(r.RankingPointsSecond),       //nolint:gosec // 点数はint32範囲内
		RankingPointsThird:    int32(r.RankingPointsThird),        //nolint:gosec // 点数はint32範囲内
		RankingPointsFour:     int32PtrFromNull(r.RankingPointsFour),
		FractionalCalculation: groupv1.FractionalCalculation(r.FractionalCalculation), //nolint:gosec // 計算方法は1-5の範囲
		UseBust:               r.UseBust,
		BustPoint:             int32PtrFromNull(r.BustPoint),
		UseChip:               r.UseChip,
		ChipPoint:             int32PtrFromNull(r.ChipPoint),
	}
}

func nullIntFromPtr(p *int32) null.Int {
	if p == nil {
		return null.Int{}
	}
	return null.IntFrom(int64(*p))
}

func int32PtrFromNull(n null.Int) *int32 {
	if !n.Valid {
		return nil
	}
	v := int32(n.Int64) //nolint:gosec // 点数はint32範囲内
	return &v
}
