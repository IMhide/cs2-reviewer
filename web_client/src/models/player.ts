import { z } from "zod";

const PlayerSchema = z.object({
  steamID: z.string().optional(),
  gameID: z.number(),
  name: z.string(),
});

type PlayerSchema = z.infer<typeof PlayerSchema>;

export default class Player {
  constructor(public readonly data: PlayerSchema) {
    Object.assign(this, data);
  }

  static fromPayload(payload: unknown) {
    const parsed = PlayerSchema.safeParse(payload);
    if (!parsed.success) {
      console.error(payload);
      throw new Error(`Invalid Player payload: ${parsed.error}`);
    }
    return new Player(parsed.data);
  }
}
