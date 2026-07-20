import { withFieldGroup } from "~/hooks/useTanstackForm";
import { chipFieldsDefaultValues } from "../GroupCreateForm/formValues";
import { useStore } from "@tanstack/react-form";
import { InputRadioGroup } from "../shared/form/InputRadioGroup";
import { USE_TYPE_OPTIONS } from "~/constants/boolOption";
import { FormControl } from "../shared/form/FormControl/FormControl";
import { InputNumber } from "../shared/form/InputNumber";
import styles from './FieldGroupChipFields.module.scss'
import { PointSuffix } from "../shared/form/suffixParts/PointSuffix";

export const FieldGroupChipFields = withFieldGroup({
  defaultValues: chipFieldsDefaultValues,
  render: function Render({ group }) {
    const useChip = useStore(group.store, (state) => state.values.useChip)
    return (
      <div role="group" aria-label="チップの設定" className={styles.chip_settings}>
        <group.Field
          name="useChip"
          children={(field) => {
            return (
              <InputRadioGroup
                name="useChip"
                items={USE_TYPE_OPTIONS}
                value={field.state.value}
                valueKey="value"
                labelKey="label"
                onChangeValue={(value) => {
                  field.handleChange(value)
                }}
                aria-label="チップの使用"
              />
            )
          }}
        />
        {useChip && (
          <group.Field
            name="chipPoint"
            children={(field) => {
              return (
                <FormControl
                  label="チップのポイント"
                  hiddenRequiredLabel
                >
                  {({ labelId, error }) => (
                    <InputNumber
                      id={labelId}
                      value={field.state.value}
                      error={error}
                      onChangeValue={(newValue) => {
                        if (newValue == null) {
                          return
                        }
                        field.handleChange
                      }}
                      suffix={<PointSuffix label="pt" />}
                    />
                  )}
                </FormControl>
              )
            }}
          />
        )}
      </div>
    )
  }
})
