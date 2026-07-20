import { FractionalCalculation, RemainderRecipient } from "~/constants/fractionalCalculation"
import { MahjongType } from "~/constants/mahjongType"

export type RankingPointFields = {
  first: number | null
  second: number | null
  third: number | null
  fourth: number | null
}

export const RankingPointFieldsDefaultValues: RankingPointFields = {
  first: 20,
  second: 10,
  third: -10,
  fourth: -20
}

export type FractionalCalculationFields = {
  /** 端数の計算方法 */
  type: FractionalCalculation
  /** 差分ポイントの受け取り */
  remainderRecipient: RemainderRecipient
}

export const FractionalCalculationFieldsDefaultValues: FractionalCalculationFields = {
  type: FractionalCalculation.DECIMAL,
  remainderRecipient: RemainderRecipient.TOP
}

export type ChipFields = {
  /** チップ */
  useChip: boolean
  /** チップのポイント */
  chipPoint: number | null
}

export const chipFieldsDefaultValues: ChipFields = {
  useChip: false,
  chipPoint: 0
}

export type GroupSettingFormValues = {
  /** グループ名 */
  groupName: string
  /** メンバー名 */
  memberNames: string[]
}

export const GroupSettingFormDefaultValues: GroupSettingFormValues = {
  groupName: '',
  memberNames: []
}

export type GroupRuleFormValues = {
  /** 麻雀の種類 */
  mahjongType: MahjongType
  /** 持ち点 */
  initialPoints: number | null
  /** 返し点 */
  returnPoints: number | null
  /** ウマ */
  rankingPoints: RankingPointFields
  /** 端数計算 */
  fractionalCalculation: FractionalCalculationFields
}

export const GroupRuleFormDefaultValues: GroupRuleFormValues = {
  mahjongType: MahjongType.THREE_PLAYER,
  initialPoints: 35,
  returnPoints: 50,
  rankingPoints: RankingPointFieldsDefaultValues,
  fractionalCalculation: FractionalCalculationFieldsDefaultValues
}

export type BustFields = {
  /** 飛び */
  useBust: boolean
  /** 飛ばし賞 */
  bustBonus: number
}

export const bustFieldsDefaultValues: BustFields = {
  useBust: false,
  bustBonus: 0
}

export type GroupDetailRuleFormValues = {
  bustSetting: BustFields,
  chipSetting: ChipFields
}

export const GroupDetailRuleFormDefaultValues: GroupDetailRuleFormValues = {
  bustSetting: bustFieldsDefaultValues,
  chipSetting: chipFieldsDefaultValues
}

export type GroupCreateFormValues = GroupSettingFormValues & GroupRuleFormValues & GroupDetailRuleFormValues

export const GroupCreateFormDefaultValues: GroupCreateFormValues = {
  ...GroupSettingFormDefaultValues,
  ...GroupRuleFormDefaultValues,
  ...GroupDetailRuleFormDefaultValues,
}
