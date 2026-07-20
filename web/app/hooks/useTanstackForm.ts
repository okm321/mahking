import { createFormHookContexts, createFormHook } from "@tanstack/react-form";
import { InputText } from "~/components/shared/form/InpuText";
import { InputNumber } from "~/components/shared/form/InputNumber";
import { InputRadioGroup } from "~/components/shared/form/InputRadioGroup";
import { InputAddToList } from "~/components/shared/form/InputTextWithButton";

const { fieldContext, formContext } =
  createFormHookContexts();

const fieldComponents = {
  InputText,
  InputNumber,
  InputAddToList,
  InputRadioGroup
}; // フィールドコンポーネントをここに定義する（例: { TextField, SelectField }）
const formComponents = {}; // フォームコンポーネントをここに定義する（例: { SubmitButton }）

export const { useAppForm, withForm, withFieldGroup } = createFormHook({
  fieldComponents,
  formComponents,
  fieldContext,
  formContext,
});
