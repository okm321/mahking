import z from "zod";

const rankingPointSchema = z.number({
  message: "半角数字を入力してください"
}).int().min(-500, {
  message: "500以下の値を入力してください"
}).max(500, {
  message: "500以上の値を入力してください"
})

export const rankingPointFieldsSchema = z.object({
  first: rankingPointSchema,
  second: rankingPointSchema,
  third: rankingPointSchema,
  fourth: z.number().int().nullable()
})

