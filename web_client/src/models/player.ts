import { z } from "zod";

type PlayerSchema = z.infer<typeof Player.PlayerSchema>;

export default class Player {
  static PlayerSchema = z.object({
    steamID: z.string(),
    gameID: z.number(),
    name: z.string(),
  });

  constructor(public readonly data: PlayerSchema) {
    Object.assign(this, data);
  }

  static fromPayload(payload: unknown) {
    const parsed = this.PlayerSchema.safeParse(payload);
    if (!parsed.success) {
      console.error(payload);
      throw new Error(`Invalid Player payload: ${parsed.error}`);
    }
    return new Player(parsed.data);
  }
}
