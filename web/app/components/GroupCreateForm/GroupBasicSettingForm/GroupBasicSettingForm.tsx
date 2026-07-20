import { FormControl } from "../../shared/form/FormControl/FormControl";
import { Heading } from "../../shared/Heading";
import { Stack } from "../../shared/Stack";
import { InputText } from "../../shared/form/InpuText";
import { withForm } from "~/hooks/useTanstackForm";
import { groupCreateFormSchema } from "~/schema/groupCreateFormSchema";
import { InputAddToList } from "../../shared/form/InputTextWithButton";
import { GroupCreateFormDefaultValues } from "../formValues";
import { revalidateLogic } from "@tanstack/react-form";

export const GroupBasicSettingForm = withForm({
  defaultValues: GroupCreateFormDefaultValues,
  validators: {
    onDynamic: groupCreateFormSchema,
  },
  validationLogic: revalidateLogic({
    mode: 'submit',
    modeAfterSubmission: 'change',
  }),
  render: function Render({ form }) {
    return (
      <section>
        <Heading as='h2'>グループ設定</Heading>
        <Stack spacing={8}>
          <form.Field
            name='groupName'
            children={(field) => {
              return (
                <FormControl
                  label="グループ名"
                  helperText="未入力だと今日の日付+メンツの名前が設定されます。"
                >
                  {({ labelId, error, ariaDescribedBy }) => (
                    <InputText
                      id={labelId}
                      name={field.name}
                      value={field.state.value}
                      aria-describedby={ariaDescribedBy}
                      placeholder="麻雀大会"
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
          <form.Field
            name='memberNames'
            mode='array'
            children={(field) => {
              const { value, meta: { errors } } = field.state
              return (
                <FormControl
                  label="メンツの名前"
                  required
                  errorMessage={errors[0]?.message}
                  helperText="同じ名前は登録できません。最大で10人まで登録できます。"
                >
                  {({ labelId, error, ariaDescribedBy }) => (
                    <InputAddToList
                      values={value}
                      name={field.name}
                      onBlur={field.handleBlur}
                      size='medium'
                      placeholder="雀太郎"
                      error={!!errors?.length || error}
                      id={labelId}
                      aria-describedby={ariaDescribedBy}
                      buttonText="追加"
                      maxCount={10}
                      onChangeValues={(values) => {
                        field.handleChange(values)
                      }}
                      onDelete={(index) => {
                        field.handleChange(field.state.value.filter((_, i) => i !== index))
                      }}
                    />
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
