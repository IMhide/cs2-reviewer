import { z } from "zod";
import { Z } from "zod-class";
import FinancialStats from "./financial_stats";
import Team from "./team";

const TeamStateSchema = Z.class({
  teamId: z.number(),
  score: z.number(),
  //financialStats: z.any(),
  financialStats: FinancialStats.schema(),
});

export default class TeamState extends TeamStateSchema {
  teamName(Teams: Team[]): string {
    return Teams.find((team) => team.id === this.teamId).name;
  }
}
