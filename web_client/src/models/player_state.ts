export default class PlayerState {
  playerGameId: number;
  playerName: string;
  isFlashed: boolean;
  positionX: number;
  positionY: number;
  positionZ: number;

  constructor(data: Object) {
    this.playerGameId = data.playerId;
    this.playerName = data.playerName;
    this.isFlashed = data.isFlashed;
    this.positionX = data.positionX;
    this.positionY = data.positionY;
    this.positionZ = data.positionZ;
  }
}
