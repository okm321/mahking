import type { ButtonHTMLAttributes, ElementType, ComponentProps, ReactNode } from 'react'
import clsx from 'clsx'
import styles from './Button.module.scss'

type ButtonSize = 's' | 'm' | 'l'
type ButtonColor = 'primary' | 'error'
type ButtonVariant = 'contained' | 'outline' | 'text'

interface BaseButtonProps {
  size?: ButtonSize
  color?: ButtonColor
  variant?: ButtonVariant
  disabled?: boolean
  bold?: boolean
  className?: string
  startIcon?: ReactNode
  endIcon?: ReactNode
  loading?: boolean
  loadingText?: string
}

type ButtonProps<T extends ElementType = 'button'> = BaseButtonProps & {
  as?: T
} & (T extends 'button' ? ButtonHTMLAttributes<HTMLButtonElement> : ComponentProps<T>)

export function Button<T extends ElementType = 'button'>({
  size = 'm',
  color = 'primary',
  variant = 'contained',
  disabled = false,
  bold = false,
  className,
  startIcon,
  endIcon,
  children,
  as,
  loading,
  loadingText,
  ...props
}: ButtonProps<T>) {
  const Component = as || 'button'

  return (
    <Component
      className={clsx(
        styles.button,
        styles[`button--${size}`],
        styles[`button--${variant}`],
        styles[`button--${color}`],
        disabled && styles['button--disabled'],
        bold && styles['button--bold'],
        loading && styles['button--loading'],
        className,
      )}
      disabled={Component === 'button' ? disabled : undefined}
      aria-disabled={disabled}
      {...props}
    >
      <span className={clsx(styles.button__content, loading && styles['button__content--hidden'])}>
        {startIcon && <span className={styles.button__startIcon}>{startIcon}</span>}
        {children}
        {endIcon && <span className={styles.button__endIcon}>{endIcon}</span>}
      </span>
      {loading && (
        <span className={styles.button__loader}>
          <img src="/pinzu1-white.svg" alt="Loading" className={styles['button__loader--spin']} />
          {loadingText && <span className={styles.button__loadingText}>{loadingText}</span>}
        </span>
      )}
    </Component>
  )
}
