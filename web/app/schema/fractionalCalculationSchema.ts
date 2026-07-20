import z from "zod";
import { FractionalCalculation, RemainderRecipient } from "~/constants/fractionalCalculation";

export const fractionalCalculationSchema = z.object({
  type: z.enum(FractionalCalculation),
  remainderRecipient: z.enum(RemainderRecipient)
})
