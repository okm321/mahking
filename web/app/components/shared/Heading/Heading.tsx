import type { PropsWithChildren } from "react"
import styles from './Heading.module.scss';
import clsx from "clsx";

type HeadingLevel = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';

type HeadingProps = PropsWithChildren<{
  /** 使用するHTML要素 */
  as?: HeadingLevel;
  /** class */
  className?: string;
}>

export function Heading({
  as,
  className,
  children
}: HeadingProps) {
  const Component = as || 'h2'
  return (
    <Component
      className={clsx(styles.heading, className)}
    >
      {children}
    </Component>
  )
}
