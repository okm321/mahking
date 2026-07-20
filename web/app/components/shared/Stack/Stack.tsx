import type { ReactNode, ElementType } from 'react'
import styles from './Stack.module.scss'

interface StackProps {
  children: ReactNode
  spacing?: number
  as?: ElementType
  className?: string
}

export function Stack({ children, spacing = 1, as: Component = 'div', className }: StackProps) {
  return (
    <Component
      className={`${styles.stack} ${className || ''}`}
      style={{ gap: `${spacing * 4}px` }}
    >
      {children}
    </Component>
  )
}
