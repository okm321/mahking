package in

import (
	"github.com/guregu/null/v6"
	"github.com/okm321/mahking/go/internal/domain"
)

type CreateGroupWithRule struct {
	Name        string   // グループ名
	MemberNames []string // メンバー名
	Rules       Rules    // ルール設定
}

type Rules struct {
	MahjongType           domain.MahjongType           // 三麻 or 四麻
	InitialPoints         int                          // 持ち点（単位: 1,000）
	ReturnPoints          int                          // 返し点（単位: 1,000）
	RankingPointsFirst    int                          // 一位のウマ
	RankingPointsSecond   int                          // 二位のウマ
	RankingPointsThird    int                          // 三位のウマ
	RankingPointsFour     null.Int                     // 四位のウマ
	FractionalCalculation domain.FractionalCalculation // 1: 切り上げ, 2: 切り捨て, 3: 四捨五入, 4: 10点未満切り上げ, 5: 10点未満切り捨て
	FractionalRecipient   domain.FractionalRecipient   // 端数を受け取る人
	UseBust               bool                         // 飛び設定
	BustPoint             null.Int                     // 飛び賞のポイント
	UseChip               bool                         // チップ設定
	ChipPoint             null.Int                     // チップのポイント
}
