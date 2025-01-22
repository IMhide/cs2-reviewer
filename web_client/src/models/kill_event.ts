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

  constructor(data: Object) {
    this.happenedAt = data.happenedAt;
    this.isHeadshot = data.isHeadshot;
    this.isWallbang = data.wallbang;
    this.isNoScope = data.noScope;
    this.isThroughSmoke = data.throughSmoke;
    this.weaponName = data.weaponName;

    this.attacker = data.attacker;
    this.victim = data.victim;
    this.assister = data.assister;
  }
}
