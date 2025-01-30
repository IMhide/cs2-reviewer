import { z } from "zod";
import { Z } from "zod-class";

const PlayerStateSchema = Z.class({
  playerGameId: z.number(),
  playerName: z.string(),
  isFlashed: z.boolean(),
  positionX: z.number(),
  positionY: z.number(),
  positionZ: z.number(),
});

export default class PlayerState extends PlayerStateSchema {}
