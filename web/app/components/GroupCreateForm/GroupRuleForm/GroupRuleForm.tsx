import { withForm } from "~/hooks/useTanstackForm";
import { groupCreateFormSchema } from "~/schema/groupCreateFormSchema";
import { Heading } from "~/components/shared/Heading";
import { Stack } from "~/components/shared/Stack";
import { FormControl } from "~/components/shared/form/FormControl/FormControl";
import { InputRadioGroup } from "~/components/shared/form/InputRadioGroup";
import { createMahjongTypeOptions, } from "~/constants/mahjongType";
import { GroupCreateFormDefaultValues } from "../formValues";
import { useMemo } from "react";
import { revalidateLogic, useStore } from "@tanstack/react-form";
import { InputNumber } from "~/components/shared/form/InputNumber";
import styles from "./GroupRuleForm.module.scss";
import { FieldGroupRankingPointFields } from "~/components/FieldGroupRankingPointFields";
import { FieldGroupFractionalCalculationFields } from "~/components/FieldGroupFractionalCalculationFields/FieldGroupFractionalCalculationFields";
import { PointSuffix } from "~/components/shared/form/suffixParts/PointSuffix";

export const GroupRuleForm = withForm({
  defaultValues: GroupCreateFormDefaultValues,
  validators: {
    onDynamic: groupCreateFormSchema,
  },
  validationLogic: revalidateLogic({
    mode: 'submit',
    modeAfterSubmission: 'change',
  }),
  render: function Render({ form }) {
    const playerNames = useStore(form.store, (state) => state.values.memberNames)
    const mahjongType = useStore(form.store, (state) => state.values.mahjongType);
    const overFourPlayers = useMemo(() => {
      return playerNames.length >= 4
    }, [playerNames.length]);

    return (
      <section>
        <Heading as='h2'>ルール設定</Heading>
        <Stack spacing={8}>
          <form.Field
            name="mahjongType"
            children={(field) => {
              const { meta: { errors } } = field.state;
              return (
                <FormControl
                  label="打ち方"
                  required
                  helperText="4人以上メンツを登録すると、四人打ちを選択できます。"
                  errorMessage={errors[0]?.message}
                >
                  {({ labelId, error, ariaDescribedBy }) => (
                    <InputRadioGroup
                      name="mahjongType"
                      aria-label="打ち方"
                      items={createMahjongTypeOptions(overFourPlayers)}
                      valueKey='value'
                      labelKey='label'
                      aria-labelledby={labelId}
                      aria-describedby={ariaDescribedBy}
                      value={field.state.value}
                      error={error}
                      onChangeValue={(value) => {
                        field.handleChange(value)
                      }}
                    />
                  )}
                </FormControl>
              )
            }}
          />
          <div className={styles.point_setting}>
            <form.Field
              name="initialPoints"
              children={(field) => {
                const { meta: { errors } } = field.state;
                return (
                  <FormControl
                    label="持ち点"
                    required
                    errorMessage={errors[0]?.message}
                    className={styles.point_form}
                  >
                    {({ labelId, error }) => (
                      <InputNumber
                        id={labelId}
                        name={field.name}
                        value={field.state.value}
                        error={error}
                        suffix={<PointSuffix label="000点" />}
                        onChangeValue={(value) => {
                          field.handleChange(value)
                        }}
                      />
                    )}
                  </FormControl>
                )
              }}
            />
            <form.Field
              name="returnPoints"
              children={(field) => {
                const { meta: { errors } } = field.state;
                return (
                  <FormControl
                    label="返し点"
                    required
                    errorMessage={errors[0]?.message}
                    className={styles.point_form}
                  >
                    {({ labelId, error }) => (
                      <InputNumber
                        id={labelId}
                        name={field.name}
                        value={field.state.value}
                        error={error}
                        suffix={<PointSuffix label="000点" />}
                        onChangeValue={(value) => {
                          field.handleChange(value)
                        }}
                      />
                    )}
                  </FormControl>
                )
              }}
            />
          </div>
          <form.Field
            name="rankingPoints"
            children={(field) => {
              const { meta: { errors } } = field.state;
              return (
                <FormControl
                  label="ウマ"
                  required
                  errorMessage={errors[0]?.message}
                >
                  {() => (
                    <FieldGroupRankingPointFields form={form} fields="rankingPoints" mahjongType={mahjongType} />
                  )}
                </FormControl>
              )
            }}
          />
          <form.Field
            name="fractionalCalculation"
            children={(field) => {
              const { meta: { errors } } = field.state;
              return (
                <FormControl
                  label="端数計算の方法"
                  required
                  helperText="小数点有効以外を選択する場合は、端数を受け取る人を選択してください。"
                  errorMessage={errors[0]?.message}
                >
                  {({ labelId, error }) => (
                    <FieldGroupFractionalCalculationFields form={form} fields="fractionalCalculation" />
                  )}
                </FormControl>
              )
            }}
          />
        </Stack>
      </section>
    )
  }
})
