import Player from "./player";

export default class Team {
  id: number;
  name: string;
  players: Player[];

  constructor(data: Object) {
    this.id = data.id;
    this.name = data.name;
    this.players = data.players.map((player) => Player.fromPayload(player));
  }
}
