export const UseType = {
  /** なし */
  NONE: 0,
  /** あり */
  USE: 1,
} as const;
export type UseType = (typeof UseType)[keyof typeof UseType];

export const USE_TYPE_LABELS: Record<UseType, string> = {
  [UseType.NONE]: "なし",
  [UseType.USE]: "あり",
}

export const USE_TYPE_OPTIONS = [
  {
    label: USE_TYPE_LABELS[UseType.NONE],
    value: false
  },
  {
    label: USE_TYPE_LABELS[UseType.USE],
    value: true
  }
]
