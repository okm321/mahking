import { withFieldGroup } from "~/hooks/useTanstackForm";
import { FractionalCalculationFieldsDefaultValues } from "../GroupCreateForm/formValues";
import { InputRadioGroup } from "../shared/form/InputRadioGroup";
import { FRACTIONAL_CALCULATION_OPTIONS, FractionalCalculation, REMAINDER_RECIPIENT_OPTIONS } from "~/constants/fractionalCalculation";
import { useStore } from "@tanstack/react-form";
import { FormControl } from "../shared/form/FormControl/FormControl";
import styles from "./FieldGroupFractionalCalculationFields.module.scss";

export const FieldGroupFractionalCalculationFields = withFieldGroup({
  defaultValues: FractionalCalculationFieldsDefaultValues,
  render: function Render({ group }) {
    const type = useStore(group.store, (state) => state.values.type)
    return (
      <div role="group" aria-label="端数計算の方法" className={styles.fractional_calculation}>
        <group.Field
          name="type"
          children={(field) => {
            return (
              <InputRadioGroup
                name="fractionalCalculationType"
                items={FRACTIONAL_CALCULATION_OPTIONS}
                value={field.state.value}
                valueKey="value"
                labelKey="label"
                onChangeValue={(value) => {
                  field.handleChange(value)
                }}
                aria-label="端数計算の方法"
              />
            )
          }}
        />
        {type !== FractionalCalculation.DECIMAL && (
          <FormControl
            label="端数を受け取る人"
            hiddenRequiredLabel
          >
            {() => {
              return (
                <group.Field
                  name="remainderRecipient"
                  children={(field) => {
                    return (
                      <InputRadioGroup
                        name="remainderRecipient"
                        items={REMAINDER_RECIPIENT_OPTIONS}
                        value={field.state.value}
                        valueKey="value"
                        labelKey="label"
                        onChangeValue={(value) => {
                          field.handleChange(value)
                        }}
                        aria-label="端数を受け取る人"
                      />
                    )
                  }}
                />
              )
            }}
          </FormControl>
        )}
      </div>
    )
  }
})
