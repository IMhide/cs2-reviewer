import { z } from "zod";

type DemoInfoSchema = z.infer<typeof DemoInfo.DemoInfoSchema>;

export default class DemoInfo {
  static DemoInfoSchema = z.object({
    mapName: z.string(),
    serverName: z.string(),
    durationInSec: z.number(),
    durationInTick: z.number(),
    durationInFrame: z.number(),
  });

  static fromPayload(payload: unknown): DemoInfo {
    const parsed = this.DemoInfoSchema.safeParse(payload);

    if (!parsed.success) {
      console.error(payload);
      throw new Error(`Invalid DemoInfo payload: ${parsed.error}`);
    }
    return new DemoInfo(parsed.data);
  }

  constructor(public readonly data: DemoInfoSchema) {
    Object.assign(this, data);
  }
}
