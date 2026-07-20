import clsx from 'clsx';
import styles from './Chip.module.scss';
import { CircleX } from 'lucide-react'
import type { AriaAttributes, ElementType } from 'react'

type ChipProps<T extends ElementType = 'div'> = {
  /** 使用するHTML要素 */
  as?: T;
  /** テキスト */
  text: string;
  /** 色 */
  color?: 'default' | 'red'
  /** 削除イベント */
  onDelete?: () => void
} & AriaAttributes

export function Chip<T extends ElementType = 'div'>({
  as,
  text,
  color = 'default',
  onDelete,
  'aria-label': ariaLabel,
  'aria-describedby': ariaDescribedBy
}: ChipProps<T>) {
  const Component = as || 'div'
  const chipLabel = ariaLabel || text
  const deleteButtonLabel = `${text}を削除`

  return (
    <Component
      className={clsx(
        styles.chip,
        color === 'red' && styles['chip-red']
      )}
      role="group"
      aria-label={chipLabel}
      aria-describedby={ariaDescribedBy}
    >
      <span className={styles.text}>{text}</span>
      {onDelete && (
        <button
          type="button"
          onClick={onDelete}
          className={styles.close_button}
          aria-label={deleteButtonLabel}
        >
          <CircleX aria-hidden="true" />
        </button>
      )}
    </Component>
  )
}
