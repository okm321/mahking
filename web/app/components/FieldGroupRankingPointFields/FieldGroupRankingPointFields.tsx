import { InputNumber } from "../shared/form/InputNumber";
import { withFieldGroup } from "~/hooks/useTanstackForm";
import { RankingPointFieldsDefaultValues } from "../GroupCreateForm/formValues";
import { FormControl } from "../shared/form/FormControl/FormControl";
import styles from "./FieldGroupRankingPointFields.module.scss"
import { MahjongType } from "~/constants/mahjongType";
import { PointSuffix } from "../shared/form/suffixParts/PointSuffix";

export const FieldGroupRankingPointFields = withFieldGroup({
  defaultValues: RankingPointFieldsDefaultValues,
  props: {
    mahjongType: MahjongType.THREE_PLAYER as MahjongType,
  },
  render: function Render({ group, mahjongType }) {
    const fieldInfos: {
      name: keyof typeof RankingPointFieldsDefaultValues;
      label: string;
    }[] = [
        { name: "first", label: "1位" },
        { name: "second", label: "2位" },
        { name: "third", label: "3位" },
      ]

    if (mahjongType === MahjongType.FOUR_PLAYER) {
      fieldInfos.push({ name: "fourth", label: "4位" });
    }
    return (
      <div role="group" aria-label="ウマ" className={styles.ranking_points}>
        {fieldInfos.map((fieldInfo) => {
          return (
            <group.Field
              name={fieldInfo.name}
              key={fieldInfo.name}
              children={(field) => {
                const { meta: { errors } } = field.state;
                return (
                  <FormControl
                    label={fieldInfo.label}
                    hiddenRequiredLabel
                    // @ts-ignore messageに推論が聞かないため
                    errorMessage={errors[0]?.message}
                  >
                    {({ labelId, error }) => (
                      <InputNumber
                        id={labelId}
                        value={field.state.value}
                        error={error}
                        onChangeValue={field.handleChange}
                        suffix={<PointSuffix label='pt' />}
                      />
                    )}
                  </FormControl>
                )
              }}
            />
          )
        })}
      </div>
    )
  }
})
