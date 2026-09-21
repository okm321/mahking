package domain

// GameRule 対局時点のルールのスナップショット。
// Rule を埋め込むことで、Rule にフィールドが増えても自動で追従する。
type GameRule struct {
	Rule
	ID     int64 // game_rules の主キー。埋め込んだ Rule.ID を隠す
	GameID int64
}

func (gr *GameRule) rankingPoint(ranking int) int {
	switch ranking {
	case 1:
		return gr.RankingPointsFirst
	case 2:
		return gr.RankingPointsSecond
	case 3:
		return gr.RankingPointsThird
	case 4:
		return int(gr.RankingPointsFour.Int64)
	default:
		return 0
	}
}

// NewGameRuleFromRule バリデーション済みの Rule からスナップショットを生成する
func NewGameRuleFromRule(rule *Rule) *GameRule {
	return &GameRule{Rule: *rule}
}
