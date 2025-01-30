import { z } from "zod";
import { Z } from "zod-class";

const FinancialStatsSchema = Z.class({
  teamId: z.number(),
  roundStartMoney: z.number(),
  freezeTimeMoney: z.number(),
  endRoundMoney: z.number(),
  spentMoney: z.number(),
});

export default class FinancialState extends FinancialStatsSchema {}
