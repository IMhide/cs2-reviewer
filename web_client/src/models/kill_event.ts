import PlayerState from "./player_state";

export default class KillEvent {
  happenedAt: number;
  isHeadshot: boolean;
  isWallbang: boolean;
  isNoScope: boolean;
  isThroughSmoke: boolean;
  weaponName: string;
  attacker: PlayerState;
  victim: PlayerState;
  assister: PlayerState;

  constructor(data: {
    happenedAt: number;
    isHeadshot: boolean;
    wallbang: boolean;
    noScope: boolean;
    throughSmoke: boolean;
    weaponName: string;
    attacker: {
      playerId: number;
      playerName: string;
      isFlashed: boolean;
      positionX: number;
      positionY: number;
      positionZ: number;
    };
    victim: {
      playerId: number;
      playerName: string;
      isFlashed: boolean;
      positionX: number;
      positionY: number;
      positionZ: number;
    };
    assister: {
      playerId: number;
      playerName: string;
      isFlashed: boolean;
      positionX: number;
      positionY: number;
      positionZ: number;
    };
  }) {
    this.happenedAt = data.happenedAt;
    this.isHeadshot = data.isHeadshot;
    this.isWallbang = data.wallbang;
    this.isNoScope = data.noScope;
    this.isThroughSmoke = data.throughSmoke;
    this.weaponName = data.weaponName;

    this.attacker = new PlayerState(data.attacker);
    this.victim = new PlayerState(data.victim);
    this.assister = new PlayerState(data.assister);
  }
}
