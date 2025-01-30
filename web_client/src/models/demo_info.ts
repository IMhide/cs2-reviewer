import { z } from "zod";
import { Z } from "zod-class";

const DemoInfoSchema = Z.class({
  mapName: z.string(),
  serverName: z.string(),
  durationInSec: z.number(),
  durationInTick: z.number(),
  durationInFrame: z.number(),
});

export default class DemoInfo extends DemoInfoSchema {}
