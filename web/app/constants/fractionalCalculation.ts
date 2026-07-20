export const FractionalCalculation = {
  /** 小数点有効 */
  DECIMAL: 0,
  /** 切り捨て */
  FLOOR: 1,
  /** 切り上げ */
  CEIL: 2,
  /** 四捨五入 */
  ROUND: 3,
  /** 五捨六入 */
  FIVE_DISCARD_SIX_ACCEPT: 4,
} as const;
export type FractionalCalculation = (typeof FractionalCalculation)[keyof typeof FractionalCalculation];

export const FRACTIONAL_CALCULATION_LABELS: Record<FractionalCalculation, string> = {
  [FractionalCalculation.DECIMAL]: "小数点有効",
  [FractionalCalculation.FLOOR]: "切り捨て",
  [FractionalCalculation.CEIL]: "切り上げ",
  [FractionalCalculation.ROUND]: "四捨五入",
  [FractionalCalculation.FIVE_DISCARD_SIX_ACCEPT]: "五捨六入",
};

export const FRACTIONAL_CALCULATION_OPTIONS = Object.entries(FractionalCalculation).map(
  ([_, value]) => ({
    value,
    label: FRACTIONAL_CALCULATION_LABELS[value],
  })
);

/** 端数受け取り人 */
export const RemainderRecipient = {
  /** トップ */
  TOP: 0,
  /** ラス */
  LAST: 1,
} as const;
export type RemainderRecipient = (typeof RemainderRecipient)[keyof typeof RemainderRecipient];

export const REMAINDER_RECIPIENT_LABELS: Record<RemainderRecipient, string> = {
  [RemainderRecipient.TOP]: "1位の人",
  [RemainderRecipient.LAST]: "最下位の人",
};

export const REMAINDER_RECIPIENT_OPTIONS = Object.entries(RemainderRecipient).map(
  ([_, value]) => ({
    value,
    label: REMAINDER_RECIPIENT_LABELS[value],
  })
);
