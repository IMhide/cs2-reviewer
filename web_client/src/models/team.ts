import Player from "./player";

export default class Team {
  id: number;
  name: string;
  players: Player[];
  dsladj;

  constructor(data: Object) {
    this.id = data.id;
    this.name = data.name;
    this.players = data.players.map((player) => new Player(player));
  }
}
