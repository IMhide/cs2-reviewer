import { z } from "zod";
import { Z } from "zod-class";

const PlayerSchema = Z.class({
  steamID: z.string(),
  gameID: z.number(),
  name: z.string(),
});

export default class Player extends PlayerSchema {}
