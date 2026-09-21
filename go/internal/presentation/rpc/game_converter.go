package rpc

import (
	"time"

	"github.com/okm321/mahking/go/internal/application/in"
	"github.com/okm321/mahking/go/internal/application/out"
	"github.com/okm321/mahking/go/internal/domain"
	gamev1 "github.com/okm321/mahking/go/internal/presentation/rpc/gen/mahking/game/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toCreateGameInput(req *gamev1.CreateGameRequest) in.CreateGame {
	scores := make([]in.GameScore, 0, len(req.GetScores()))
	for _, s := range req.GetScores() {
		scores = append(scores, in.GameScore{
			MemberID:  s.GetMemberId(),
			Seat:      domain.Seat(s.GetSeat()),
			Ranking:   int(s.GetRanking()),
			RawScore:  int(s.GetRawScore()),
			ChipCount: nullIntFromPtr(s.ChipCount),
			IsBusted:  s.GetIsBusted(),
		})
	}

	var playedAt time.Time
	if req.PlayedAt != nil {
		playedAt = req.GetPlayedAt().AsTime()
	}

	return in.CreateGame{
		GroupUID: req.GetGroupUid(),
		Note:     req.GetNote(),
		PlayedAt: playedAt,
		Scores:   scores,
	}
}

func toProtoGame(g *out.Game) *gamev1.Game {
	scores := make([]*gamev1.GameScore, 0, len(g.Scores))
	for _, s := range g.Scores {
		scores = append(scores, &gamev1.GameScore{
			Id:        s.ID,
			MemberId:  s.MemberID,
			Seat:      gamev1.Seat(s.Seat), //nolint:gosec // 席は1-4の範囲
			Ranking:   int32(s.Ranking),    //nolint:gosec // 順位は1-4の範囲
			RawScore:  int32(s.RawScore),   //nolint:gosec // 素点はint32範囲内
			Point:     s.Point,
			ChipCount: int32PtrFromNull(s.ChipCount),
			IsBusted:  s.IsBusted,
		})
	}
	return &gamev1.Game{
		Id:       g.ID,
		GroupId:  g.GroupID,
		Note:     g.Note,
		PlayedAt: timestamppb.New(g.PlayedAt),
		Rules:    toProtoRules(g.Rule),
		Scores:   scores,
	}
}
