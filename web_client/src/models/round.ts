import { z } from "zod";
import { Z } from "zod-class";
import KillEvent from "./kill_event";
import TeamState from "./team_state";
import Team from "./team";

const RoundSchema = Z.class({
  roundNumber: z.number(),
  startTime: z.number(),
  endTime: z.number(),
  winningReason: z.number(),
  winnerState: TeamState.schema(),
  loserState: TeamState.schema(),
  killFeed: z.array(KillEvent.schema()),
});

export default class Round extends RoundSchema {
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
