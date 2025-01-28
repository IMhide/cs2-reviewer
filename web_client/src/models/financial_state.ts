import { z } from "zod";

type FinancialStateSchema = z.infer<typeof FinancialState.FinancialStateSchema>;

export default class FinancialState {
  static FinancialStateSchema = z.object({
    teamId: z.number(),
    RoundStartMoney: z.number(),
    FreezeTimeEndMoney: z.number(),
    EndRoundMoney: z.number(),
    TotalSpentMoney: z.number(),
  });

  static fromPayload(payload: unknown) {
    const parsed = this.FinancialStateSchema.safeParse(payload);
    if (!parsed.success) {
      console.error(payload);
      throw new Error(`Invalid FinancialState payload: ${parsed.error}`);
    }
    return new FinancialState(parsed.data);
  }
  constructor(public readonly data: FinancialStateSchema) {
    Object.assign(this, data);
  }
}
