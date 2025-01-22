import FinancialState from "./financial_state";
import Team from "./team";

export default class TeamState {
  teamId: number;
  score: number;
  financialState: FinancialState;

  constructor(data: Object) {
    this.teamId = data.id;
    this.score = data.score;
    this.financialState = new FinancialState(data.financialStats);
  }

  teamName(Teams: Team[]): string {
    return Teams.find((team) => team.id === this.teamId).name;
  }
}
