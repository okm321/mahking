import type { FC } from "react";
import styles from "./PointSuffix.module.scss"

type PointSuffixProps = {
  label: string;
}

export const PointSuffix: FC<PointSuffixProps> = ({
  label
}) => {
  return <span className={styles.point_suffix}>{label}</span>
}
