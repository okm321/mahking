import { useAppForm } from "~/hooks/useTanstackForm"
import { groupCreateFormSchema } from "~/schema/groupCreateFormSchema"
import { Stack } from "../shared/Stack"
import { Users } from "lucide-react";
import { GroupBasicSettingForm } from "./GroupBasicSettingForm"
import { Button } from "../shared/Button"
import styles from "./GroupCreateForm.module.scss"
import { GroupRuleForm } from "./GroupRuleForm";
import { GroupCreateFormDefaultValues } from "./formValues";
import { revalidateLogic } from "@tanstack/react-form";
import { GroupRuleDetailForm } from "./GroupRuleDetailForm";

export function GroupCreateForm() {
  const form = useAppForm({
    defaultValues: GroupCreateFormDefaultValues,
    validators: {
      // onSubmit: groupCreateFormSchema,
      // onChangeAsync: groupCreateFormSchema,
      // onChangeAsyncDebounceMs: 500,
      onDynamic: groupCreateFormSchema,
    },
    validationLogic: revalidateLogic({
      mode: 'submit',
      modeAfterSubmission: 'change',
    }),
    onSubmit: async ({ value }) => {
      console.log("グループ作成リクエスト:", value)
      await new Promise((resolve) => setTimeout(resolve, 2000)) // Simulate API call
    },
  })

  return (
    <form onSubmit={(e) => {
      e.preventDefault()
      e.stopPropagation()
      form.handleSubmit()
    }} className={styles.group_create_form}>
      <Stack spacing={8}>
        <GroupBasicSettingForm form={form} />
        <GroupRuleForm form={form} />
        <GroupRuleDetailForm form={form} />
      </Stack>
      <form.Subscribe
        selector={(state) => [state.isSubmitting]}
        children={([isSubmitting]) => (
          <Button
            size='l'
            type="submit"
            startIcon={<Users />}
            loading={isSubmitting}
            loadingText="グループ作成中"
            className={styles.submit_button}
          >グループを作成する</Button>
        )}
      />
    </form>
  )
}
