package domain

import (
	"context"
	"time"

	pkgerror "github.com/okm321/mahking/go/pkg/error"
	pkgtime "github.com/okm321/mahking/go/pkg/time"
)

type Game struct {
	ID         int64
	GroupID    int64
	Note       string       // 対局メモ（最大200文字）
	PlayedAt   time.Time    // 対局日時
	GameRule   *GameRule    // 対局時ルールスナップショット
	GameScores []*GameScore // 各プレイヤーの成績
}

const MaxGameNoteLength = 200

type NewGameArgs struct {
	Note       string
	PlayedAt   time.Time
	GameRule   *GameRule
	GameScores []*GameScore
}

func NewGame(ctx context.Context, groupID int64, args NewGameArgs) (*Game, error) {
	playedAt := args.PlayedAt
	if playedAt.IsZero() {
		playedAt = pkgtime.Now(ctx)
	}

	g := &Game{
		GroupID:    groupID,
		Note:       args.Note,
		PlayedAt:   playedAt,
		GameRule:   args.GameRule,
		GameScores: args.GameScores,
	}

	if err := g.Validate(); err != nil {
		return nil, err
	}

	g.calculatePoints()

	return g, nil
}

func (g *Game) Validate() error {
	if g.GameRule == nil {
		return pkgerror.NewClientError("対局ルールは必須です")
	}
	if err := optionalText("対局メモ", g.Note, MaxGameNoteLength); err != nil {
		return err
	}
	requiredCount := g.GameRule.MahjongType.RequiredMemberCount()
	if len(g.GameScores) != requiredCount {
		return pkgerror.NewClientErrorf(
			"%sのスコア数は%d人分必要です: %d人分",
			g.GameRule.MahjongType.String(),
			requiredCount,
			len(g.GameScores),
		)
	}

	seatMap := make(map[Seat]bool)
	memberIDMap := make(map[int64]bool)
	rankingMap := make(map[int]bool)
	for _, score := range g.GameScores {
		if !score.Seat.IsValid() {
			return pkgerror.NewClientErrorf("席が不正です: %d", score.Seat)
		}
		if score.Ranking < 1 || score.Ranking > len(g.GameScores) {
			return pkgerror.NewClientErrorf("順位は1〜%dの範囲で指定してください: %d", len(g.GameScores), score.Ranking)
		}
		if seatMap[score.Seat] {
			return pkgerror.NewClientErrorf("席が重複しています: %s", score.Seat)
		}
		seatMap[score.Seat] = true

		if memberIDMap[score.MemberID] {
			return pkgerror.NewClientErrorf("メンバーが重複しています: メンバーID %d", score.MemberID)
		}
		memberIDMap[score.MemberID] = true

		if rankingMap[score.Ranking] {
			return pkgerror.NewClientErrorf("順位が重複しています: %d位", score.Ranking)
		}
		rankingMap[score.Ranking] = true

		if g.GameRule.UseChip && !score.ChipCount.Valid {
			return pkgerror.NewClientErrorf("チップ枚数は必須です: メンバーID %d", score.MemberID)
		}
		if !g.GameRule.UseChip && score.ChipCount.Valid {
			return pkgerror.NewClientErrorf("チップ設定が無効な場合、チップ枚数は指定できません: メンバーID %d", score.MemberID)
		}

		if !g.GameRule.UseBust && score.IsBusted {
			return pkgerror.NewClientErrorf("飛び設定が無効な場合、飛びフラグは指定できません: メンバーID %d", score.MemberID)
		}
	}

	return nil
}

func (g *Game) calculatePoints() {
	playerCount := len(g.GameScores)
	oka := (g.GameRule.ReturnPoints - g.GameRule.InitialPoints) * playerCount

	for _, score := range g.GameScores {
		diff := float64(score.RawScore-g.GameRule.ReturnPoints*10) / 10
		point := g.GameRule.FractionalCalculation.Apply(diff) + float64(g.GameRule.rankingPoint(score.Ranking))
		if score.Ranking == 1 {
			point += float64(oka)
		}
		score.Point = point
	}

	// 小数点有効の場合は端数調整不要
	if g.GameRule.FractionalCalculation.IsDecimal() {
		return
	}

	// 丸めで合計が0にならない場合、端数を受け取る人に調整分を加算
	var total float64
	for _, score := range g.GameScores {
		total += score.Point
	}
	remainder := -total // 合計を0にするための調整値
	if remainder == 0 {
		return
	}

	for _, score := range g.GameScores {
		if score.Ranking == g.GameRule.FractionalRecipient.TargetRanking(playerCount) {
			score.Point += remainder
			break
		}
	}
}

// GameRepository 永続化層のインタフェース
type GameRepository interface {
	Create(ctx context.Context, game *Game) (*Game, error)
	List(ctx context.Context, groupID int64) ([]Game, error)
}
