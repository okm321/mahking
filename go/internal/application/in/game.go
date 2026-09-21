package in

import (
	"time"

	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/domain"
)

type CreateGame struct {
	GroupUID string
	Note     string
	PlayedAt time.Time // ゼロ値なら現在時刻
	Scores   []GameScore
}

type GameScore struct {
	MemberID  int64
	Seat      domain.Seat
	Ranking   int
	RawScore  int
	ChipCount null.Int
	IsBusted  bool
}
