import { withFieldGroup } from "~/hooks/useTanstackForm";
import { bustFieldsDefaultValues } from "../GroupCreateForm/formValues";
import { InputRadioGroup } from "../shared/form/InputRadioGroup";
import { USE_TYPE_OPTIONS } from "~/constants/boolOption";
import { useStore } from "@tanstack/react-form";
import { InputNumber } from "../shared/form/InputNumber";
import { FormControl } from "../shared/form/FormControl/FormControl";
import { PointSuffix } from "../shared/form/suffixParts/PointSuffix";
import styles from "./FieldGroupBustFields.module.scss";

export const FieldGroupBustFields = withFieldGroup({
  defaultValues: bustFieldsDefaultValues,
  render: function Render({ group }) {
    const useBust = useStore(group.store, (state) => state.values.useBust);
    return (
      <div role="group" aria-label="飛び設定" className={styles.bust_settings}>
        <group.Field
          name="useBust"
          children={(field) => {
            return (
              <InputRadioGroup
                name="useBust"
                items={USE_TYPE_OPTIONS}
                value={field.state.value}
                valueKey="value"
                labelKey="label"
                onChangeValue={(value) => {
                  field.handleChange(value)
                }}
                aria-label="飛び設定"
              />
            )
          }}
        />
        {useBust && (
          <group.Field
            name="bustBonus"
            children={(field) => {
              return (
                <FormControl
                  label="飛ばし賞"
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
                      }
                      }
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
