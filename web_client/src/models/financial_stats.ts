import { z } from "zod";
import { Z } from "zod-class";

const FinancialStatsSchema = Z.class({
  teamId: z.number(),
  RoundStartMoney: z.number(),
  FreezeTimeEndMoney: z.number(),
  EndRoundMoney: z.number(),
  TotalSpentMoney: z.number(),
});

export default class FinancialState extends FinancialStatsSchema {}
