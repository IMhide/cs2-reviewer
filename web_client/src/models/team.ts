import { z } from "zod";
import Player from "./player";

type TeamSchema = z.infer<typeof Team.TeamSchema>;

export default class Team {
  static TeamSchema = z.object({
    id: z.number(),
    name: z.string(),
    players: Player.PlayerSchema.array(),
  });

  constructor(public readonly data: TeamSchema) {
    Object.assign(this, data);
  }

  static fromPayload(payload: unknown): Team {
    const parsed = this.TeamSchema.safeParse(payload);
    if (!parsed.success) {
      console.error(payload);
      throw new Error(`Invalid Team payload: ${parsed.error}`);
    }
    return new Team(parsed.data);
  }
}
