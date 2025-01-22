import KillEvent from "./kill_event";
import TeamState from "./team_state";
import Team from "./team";

export default class Round {
  roundNumber: number;
  startTime: number;
  endTime: number;
  winningReason: number;
  winnerState: TeamState;
  loserState: TeamState;
  killFeed: killEvent[];

  constructor(data: Object) {
    this.roundNumber = data.roundNumber;
    this.startTime = data.startTime;
    this.endTime = data.endTime;
    this.winningReason = data.winningReason;
    this.winnerState = new TeamState(data.winnerState);
    this.loserState = new TeamState(data.loserState);
    this.killFeed = data.killFeed.map((killEvent) => new KillEvent(killEvent));
  }

  teamsScore(teams: Team[]): Object[] {
    return [
      {
        teamId: this.winnerState.teamId,
        teamName: this.winnerState.teamName(teams),
        score: this.winnerState.score,
      },
      {
        teamId: this.loserState.teamId,
        teamName: this.loserState.teamName(teams),
        score: this.loserState.score,
      },
    ].sort((a, b) => b.teamId < a.teamId);
  }

  teamScoreToString(teams: Team[]): string {
    const scores = this.teamsScore(teams);

    return `${scores[0].teamName} ${scores[0].score} - ${scores[1].score}  ${scores[1].teamName}`;
  }
}
