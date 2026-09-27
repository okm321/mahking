package out

import (
	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/domain"
)

// Group is a view model returned by usecases.
type Group struct {
	ID      int64
	UID     string
	Name    string
	Members []Member
	Rule    Rule
}

type Member struct {
	ID   int64
	Name string
}

type Rule struct {
	MahjongType           domain.MahjongType
	InitialPoints         int
	ReturnPoints          int
	RankingPointsFirst    int
	RankingPointsSecond   int
	RankingPointsThird    int
	RankingPointsFour     null.Int
	FractionalCalculation domain.FractionalCalculation
	FractionalRecipient   domain.FractionalRecipient
	UseBust               bool
	BustPoint             null.Int
	UseChip               bool
	ChipPoint             null.Int
}

func NewGroup(g domain.Group) Group {
	members := make([]Member, 0, len(g.Members))
	for _, m := range g.Members {
		members = append(members, Member{ID: m.ID, Name: m.Name})
	}
	out := Group{
		ID:      g.ID,
		UID:     g.UID,
		Name:    g.Name,
		Members: members,
	}
	if g.Rule != nil {
		out.Rule = NewRule(*g.Rule)
	}
	return out
}

func NewRule(r domain.Rule) Rule {
	return Rule{
		MahjongType:           r.MahjongType,
		InitialPoints:         r.InitialPoints,
		ReturnPoints:          r.ReturnPoints,
		RankingPointsFirst:    r.RankingPointsFirst,
		RankingPointsSecond:   r.RankingPointsSecond,
		RankingPointsThird:    r.RankingPointsThird,
		RankingPointsFour:     r.RankingPointsFour,
		FractionalCalculation: r.FractionalCalculation,
		FractionalRecipient:   r.FractionalRecipient,
		UseBust:               r.UseBust,
		BustPoint:             r.BustPoint,
		UseChip:               r.UseChip,
		ChipPoint:             r.ChipPoint,
	}
}
