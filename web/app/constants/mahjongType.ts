export const MahjongType = {
  THREE_PLAYER: 0,
  FOUR_PLAYER: 1,
} as const
export type MahjongType = (typeof MahjongType)[keyof typeof MahjongType]

export const MAHJONG_TYPE_LABELS: Record<MahjongType, string> = {
  [MahjongType.THREE_PLAYER]: "三人打ち",
  [MahjongType.FOUR_PLAYER]: "四人打ち",
}

export const MAHJONG_TYPE_OPTIONS = Object.entries(MahjongType).map(
  ([_, value]) => ({
    value,
    label: MAHJONG_TYPE_LABELS[value],
  })
)

/** 
 * 麻雀の種類を選択するオプションを生成します。
 * @param overFourPlayers - 4人以上のプレイヤーがいるか
 */
export const createMahjongTypeOptions = (overFourPlayers: boolean) => {
  return [
    {
      label: MAHJONG_TYPE_LABELS[MahjongType.THREE_PLAYER],
      value: MahjongType.THREE_PLAYER
    },
    {
      label: MAHJONG_TYPE_LABELS[MahjongType.FOUR_PLAYER],
      value: MahjongType.FOUR_PLAYER,
      disabled: !overFourPlayers
    }
  ]
}
