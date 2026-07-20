import { useId, type AriaAttributes } from "react"
import styles from "./InputRadioGroup.module.scss"
import clsx from "clsx";

type InputRadioGroupProps<
  Item extends Record<string, any>,
  ValueKey extends keyof Item,
> = {
  /** value */
  value: Item[ValueKey];
  /** name */
  name: string;
  /** 項目 */
  items: Item[];
  /** valueのキー */
  valueKey: ValueKey;
  /** labelのキー */
  labelKey: keyof Item;
  /** エラーか */
  error?: boolean;
  /** 値変更時のコールバック */
  onChangeValue: (value: Item[ValueKey]) => void;
} & AriaAttributes

export function InputRadioGroup<
  Item extends Record<string, any>,
  ValueKey extends keyof Item
>({
  value,
  name,
  items,
  valueKey,
  labelKey,
  error,
  onChangeValue,
  "aria-label": ariaLabel,
  "aria-describedby": ariaDescribedBy,
}: InputRadioGroupProps<Item, ValueKey>) {
  return (
    <div
      role="radiogroup"
      aria-label={ariaLabel}
      aria-describedby={ariaDescribedBy}
      aria-invalid={error}
      className={clsx(
        styles.input_radio_group,
        error && styles['input_radio_group-error']
      )}
    >
      {items.map((item) => {
        const itemId = `radio-item-${name}-${item[valueKey]}`
        return (
          <span key={itemId} className={styles.input_radio_item}>
            <span className={styles.input_radio_wrapper}>
              <input
                id={itemId}
                type='radio'
                name={name}
                aria-invalid={error}
                checked={value === item[valueKey]}
                disabled={item.disabled}
                value={item[valueKey]}
                onChange={() => onChangeValue(item[valueKey])}
                className={styles.input_radio}
              />
              <span className={styles.radio_button} aria-hidden="true" />
            </span>
            <label
              htmlFor={itemId}
              className={styles.input_label}
              data-disabled={item.disabled ? "true" : undefined}
            >{item[labelKey]}</label>
          </span>
        )
      })}
    </div>
  )
}
