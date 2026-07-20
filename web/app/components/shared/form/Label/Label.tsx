import clsx from "clsx";
import type { FC, PropsWithChildren } from "react";
import styles from "./Label.module.scss"
import { RequiredChip } from "../RequiredChip";

type LabelProps = PropsWithChildren<{
  /** htmlFor */
  htmlFor: string
  /** 必須かどうか */
  required?: boolean
  /** 必須チップを隠すか */
  hiddenChip?: boolean
}>

export function Label({ htmlFor, required, hiddenChip, children }: LabelProps) {
  return (
    <label
      htmlFor={htmlFor}
      className={clsx(styles.label)}
    ><>
        {!hiddenChip && <RequiredChip required={required} />}
        {children}
      </></label>
  )
}
