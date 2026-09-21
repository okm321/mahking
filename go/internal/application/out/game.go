package out

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/domain"
)

type Game struct {
	ID       int64
	GroupID  int64
	Note     string
	PlayedAt time.Time
	Rule     Rule
	Scores   []GameScore
}

type GameScore struct {
	ID        int64
	MemberID  int64
	Seat      domain.Seat
	Ranking   int
	RawScore  int
	Point     float64
	ChipCount null.Int
	IsBusted  bool
}

func NewGame(g domain.Game) Game {
	scores := make([]GameScore, 0, len(g.GameScores))
	for _, s := range g.GameScores {
		scores = append(scores, GameScore{
			ID:        s.ID,
			MemberID:  s.MemberID,
			Seat:      s.Seat,
			Ranking:   s.Ranking,
			RawScore:  s.RawScore,
			Point:     s.Point,
			ChipCount: s.ChipCount,
			IsBusted:  s.IsBusted,
		})
	}
	return Game{
		ID:       g.ID,
		GroupID:  g.GroupID,
		Note:     g.Note,
		PlayedAt: g.PlayedAt,
		Rule:     NewRule(g.GameRule.Rule),
		Scores:   scores,
	}
}
