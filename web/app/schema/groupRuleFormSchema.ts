import z from "zod";
import { MahjongType } from "~/constants/mahjongType";
import { rankingPointFieldsSchema } from "./rankingPointSchema";
import { fractionalCalculationSchema } from "./fractionalCalculationSchema";

const pointsSchema = z.number({
  message: "半角数字を入力してください"
}).int().min(10, {
  message: "10以上の値を入力してください"
}).max(50, {
  message: "50以下の値を入力してください"
})

export const groupRuleFormSchema = z.object({
  mahjongType: z.enum(MahjongType),
  initialPoints: pointsSchema,
  returnPoints: pointsSchema,
  rankingPoints: rankingPointFieldsSchema,
  fractionalCalculation: fractionalCalculationSchema,
})
