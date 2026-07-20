import z from "zod";

export const groupSettingFormSchema = z.object({
  groupName: z.string().max(50, {
    message: "50文字以内で入力してください"
  }),
  memberNames: z.array(z.string().max(10)).min(3, {
    message: "メンツは3人以上登録してください"
  }).max(10, {
    message: "メンツは最大10人まで登録できます"
  }).check((ctx) => {
    const uniqueNames = ctx.value.filter((name, i, self) => {
      return self.indexOf(name) === i && i !== self.lastIndexOf(name)
    })
    if (uniqueNames.length > 0) {
      ctx.issues.push({
        code: "custom",
        message: `同じ名前は登録できません: ${Array.from(uniqueNames).join(", ")}`,
        input: ctx.value,
      })
    }
  }),
})
