export default class Player {
  steamId: number;
  gameId: number;
  name: string;

  constructor(data: Object) {
    this.steamId = data.steamID;
    this.gameId = data.gameID;
    this.name = data.name;
  }
}
