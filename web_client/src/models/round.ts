import killEvent from './kill_event';
import TeamState from './team_state';

export default class Round {
  roundNumber: number;
  startTime: number;
  endTime: number;
  winningReason: number;
  winnerState: TeamState;
  looserState: TeamState;
  killFeed: killEvent[];

  constructor(data: Object) {
    this.roundNumber = data.roundNumber;
    this.startTime = data.startTime;
    this.endTime = data.endTime;
    this.winningReason = data.winningReason;
    this.winnerState = new TeamState(data.winnerState);
    this.looserState = new TeamState(data.looserState);
    this.killFeed = data.killFeed.map(killEvent => new KillEvent(killEvent));
  }
}
