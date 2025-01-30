import { z } from "zod";
import { Z } from "zod-class";
import Player from "./player";

const TeamSchema = Z.class({
  id: z.number(),
  name: z.string(),
  players: z.array(Player.schema()),
});

export default class Team extends TeamSchema {}
