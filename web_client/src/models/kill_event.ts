import { z } from "zod";
import { Z } from "zod-class";
import PlayerState from "./player_state";

const KillEventSchema = Z.class({
  happenedAt: z.number(),
  isHeadshot: z.boolean(),
  isWallbang: z.boolean(),
  isNoScope: z.boolean(),
  isThroughSmoke: z.boolean(),
  weaponName: z.string(),
  attacker: PlayerState.schema(),
  victim: PlayerState.schema(),
  assister: PlayerState.schema(),
});

export default class KillEvent extends KillEventSchema {}
