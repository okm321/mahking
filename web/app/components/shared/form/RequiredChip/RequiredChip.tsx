import clsx from "clsx";
import styles from "./RequiredChip.module.scss"

type RequiredChipProps = {
  /** 必須かどうか */
  required?: boolean;
}

export function RequiredChip({ required }: RequiredChipProps) {
  return (
    <span
      className={clsx(
        styles.chip,
        required ? styles.chip__required : styles.chip__optional
      )}
    >{required ? '必須' : '任意'}</span>
  )
}
