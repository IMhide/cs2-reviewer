import FinancialState from './financial_state';

export default class TeamState {
  teamId: number;
  score: number;
  financialState: FinancialState;

  constructor(data: Object) {
    this.teamId = data.id;
    this.score = data.score;

//    this.financialState = new FinancialState(data.financialState);
  }
}
