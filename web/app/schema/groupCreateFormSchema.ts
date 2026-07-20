import z from "zod";
import { MahjongType } from "~/constants/mahjongType";
import { groupSettingFormSchema } from "./groupSettingFormSchema";
import { groupRuleFormSchema } from "./groupRuleFormSchema";
import { groupRuleDetailFormSchema } from "./groupRuleDetailFormSchema";

export const groupCreateFormSchema = z.object({
  ...groupSettingFormSchema.shape,
  ...groupRuleFormSchema.shape,
  ...groupRuleDetailFormSchema.shape,
}).refine((data) => {
  if (data.memberNames.length < 4 && data.mahjongType === MahjongType.FOUR_PLAYER) {
    return false;
  }
  return true;
}, {
  message: "4人未満のメンツでは四人打ちを選択できません",
  path: ["mahjongType"],
}).refine(({ initialPoints, returnPoints }) => {
  if (initialPoints == null || returnPoints == null) {
    return true;
  }
  if (initialPoints > returnPoints) {
    return false;
  }

  return true
}, {
  message: "返し点は持ち点以上にしてください",
  path: ["returnPoints"],
}).refine(({ mahjongType, rankingPoints }) => {
  if (mahjongType === MahjongType.FOUR_PLAYER && rankingPoints.fourth == null) {
    return false;
  }
  return true
}, {
  message: "半角数字を入力してください",
  path: ["rankingPoints.fourth"]
})
  .refine(({ mahjongType, rankingPoints }) => {
    if (mahjongType === MahjongType.THREE_PLAYER) {
      const sum = rankingPoints.first + rankingPoints.second + rankingPoints.third;
      if (sum !== 0) {
        return false;
      }
      return true;
    }
    const sum = rankingPoints.first + rankingPoints.second + rankingPoints.third + (rankingPoints.fourth ?? 0);
    if (sum !== 0) {
      return false;
    }
    return true
  }, {
    message: "ウマの合計は0になるようにしてください",
    path: ["rankingPoints"],
  })

export type GroupCreateFormType = z.infer<typeof groupCreateFormSchema>;

