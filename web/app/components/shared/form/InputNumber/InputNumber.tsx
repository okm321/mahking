import { useState, useEffect } from "react"
import type { AriaAttributes, FocusEvent, ReactNode } from "react"
import styles from "../InpuText/InputText.module.scss"
import clsx from "clsx"

type InputNumberProps = {
  /** id */
  id?: string
  /** name */
  name?: string
  /** 値 */
  value: number | null
  onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void
  /** 値変更のコールバック */
  onChangeValue: (value: number | null) => void
  /** サイズ */
  size?: 'small' | 'medium'
  /** 非活性か */
  disabled?: boolean
  /** プレースホルダー */
  placeholder?: string
  /** エラーか */
  error?: boolean
  /** input-mode */
  inputMode?: React.HTMLAttributes<HTMLInputElement>['inputMode']
  /** suffix */
  suffix?: ReactNode
} & AriaAttributes

export function InputNumber({
  id,
  name,
  value,
  onBlur,
  onChangeValue,
  size = 'medium',
  disabled,
  placeholder,
  error,
  inputMode = 'numeric',
  suffix,
  "aria-describedby": ariaDescribedBy
}: InputNumberProps) {
  const [localValue, setLocalValue] = useState<string>('')

  const handleInputFocus = (e: FocusEvent<HTMLInputElement, Element>) => {
    setTimeout(() => {
      e.target.scrollIntoView({
        behavior: 'smooth',
        block: 'center',
        inline: 'nearest'
      });
    }, 300);
  };

  useEffect(() => {
    setLocalValue(value === null ? '' : String(value))
  }, [value])

  const convertAndValidateNumber = (inputValue: string): { value: number | null; displayValue: string } => {
    const converted = convertZenkakuToHankaku(inputValue)

    if (converted === '') {
      return { value: null, displayValue: '' }
    }

    // 整数のみを許可（マイナス記号と数値のみ）
    if (!/^-?\d+$/.test(converted)) {
      // 無効な値の場合は空文字列を表示値として返す
      return { value: null, displayValue: '' }
    }

    const numericValue = Number(converted)
    if (!isNaN(numericValue) && Number.isInteger(numericValue)) {
      return { value: numericValue, displayValue: converted }
    }

    return { value: null, displayValue: '' }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const inputValue = e.target.value
    setLocalValue(inputValue)
    // onChange時は値の妥当性をチェックせず、入力をそのまま反映
  }

  const handleBlur = () => {
    const { value, displayValue } = convertAndValidateNumber(localValue)
    setLocalValue(displayValue)
    onChangeValue(value)
  }

  return (
    <div className={styles.input_text_wrapper}>
      <input
        id={id}
        name={name}
        value={localValue}
        onChange={handleChange}
        onBlur={(e) => {
          handleBlur()
          onBlur?.(e)
        }}
        onFocus={handleInputFocus}
        inputMode={inputMode}
        disabled={disabled}
        placeholder={placeholder}
        aria-invalid={error}
        aria-describedby={ariaDescribedBy}
        autoComplete="off"
        autoCorrect="off"
        className={clsx(
          styles.input_text,
          size === 'medium' && styles.input_text__medium,
          error && styles['input_text-error'],
          disabled && styles['input_text-disabled']
        )}
      />
      {suffix}
    </div>
  )
}

const convertZenkakuToHankaku = (str: string): string => {
  return str.replace(/[０-９]/g, (s) => {
    return String.fromCharCode(s.charCodeAt(0) - 0xFEE0)
  })
}

