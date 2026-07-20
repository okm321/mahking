import { useMemo, useRef, useState } from 'react'
import clsx from 'clsx'
import styles from './InputAddToList.module.scss'
import type { AriaAttributes, FocusEvent } from 'react'
import { Chip } from '../../Chip'

type InputAddToListProp = {
  /** id */
  id?: string
  /** name */
  name?: string
  /** value */
  values: string[]
  onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void
  /** 値変更時 */
  onChangeValues: (value: string[]) => void
  /** 値が正しいか */
  onCheckValidValue?: (value: string) => boolean
  /** サイズ */
  size?: 'small' | 'medium'
  /** 非活性か */
  disabled?: boolean
  /** プレースホルダー */
  placeholder?: string
  /** エラーか */
  error?: boolean
  /** ボタンのテキスト */
  buttonText?: string
  /** 最大数 */
  maxCount?: number
  /** 削除イベント */
  onDelete?: (index: number) => void
} & AriaAttributes

export function InputAddToList({
  id,
  name,
  values,
  onBlur,
  onChangeValues,
  onCheckValidValue,
  size = 'medium',
  disabled,
  placeholder,
  error,
  buttonText = '送信',
  maxCount,
  onDelete,
  "aria-describedby": ariaDescribedBy
}: InputAddToListProp) {
  const [inputValue, setInputValue] = useState<string>('')

  const handleInputFocus = (e: FocusEvent<HTMLInputElement, Element>) => {
    setTimeout(() => {
      e.target.scrollIntoView({
        behavior: 'smooth',
        block: 'center',
        inline: 'nearest'
      });
    }, 300);
  };


  const handleSubmit = () => {
    if (maxCount && values.length >= maxCount) {
      return
    }
    if (!disabled && inputValue.trim()) {
      if (onCheckValidValue != null && !onCheckValidValue(inputValue.trim())) {
        return
      }
      onChangeValues([...values, inputValue.trim()])
      setInputValue('')
    }
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && !e.nativeEvent.isComposing) {
      e.stopPropagation()
      e.preventDefault()
      handleSubmit()
    }
  }

  const disabledSubmitButton = useMemo(() => {
    return disabled || (maxCount != null && values.length >= maxCount) || !inputValue.trim()
  }, [values, disabled, maxCount, inputValue])

  const uniqueValues = useMemo(() => {
    return values.filter((name, i, self) => {
      return self.indexOf(name) === i && i !== self.lastIndexOf(name)
    })
  }, [values])

  return (
    <>
      <div
        className={clsx(
          styles.input_text_with_button,
          size === 'medium' ? styles.input_text_with_button__medium : styles.input_text_with_button__small,
          error && styles['input_text_with_button-error'],
          disabled && styles['input_text_with_button-disabled']
        )}
      >
        <input
          id={id}
          name={name}
          type="text"
          onFocus={handleInputFocus}
          className={styles.input_text}
          placeholder={placeholder}
          disabled={disabled}
          onBlur={onBlur}
          value={inputValue}
          onChange={(event) => setInputValue(event.target.value)}
          onKeyDown={handleKeyDown}
          aria-invalid={error}
          aria-describedby={ariaDescribedBy}
          enterKeyHint="done"
        />
        <button
          type="button"
          className={clsx(
            styles.button,
            disabledSubmitButton && styles[`button-disabled`]
          )}
          disabled={disabledSubmitButton}
          onClick={handleSubmit}
        >
          {buttonText}
        </button>
      </div>

      {values.length > 0 && (
        <ul className={styles.list}>
          {values.map((member, index) => (
            <Chip
              as="li"
              color={uniqueValues.includes(member) ? 'red' : 'default'}
              key={index}
              text={member}
              onDelete={onDelete ? () => onDelete(index) : undefined}
            />
          ))}
        </ul>
      )}
    </>
  )
}
