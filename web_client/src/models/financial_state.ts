export default class FinancialState {
  teamId: number;
  RoundStartMoney: number;
  FreezeTimeEndMoney: number;
  EndRoundMoney: number;
  TotalSpentMoney: number;

  constructor(data: Object) {
    this.teamId = data.teamId;
    this.RoundStartMoney = data.RoundStartMoney;
    this.FreezeTimeEndMoney = data.FreezeTimeEndMoney;
    this.EndRoundMoney = data.EndRoundMoney;
    this.TotalSpentMoney = data.TotalSpentMoney;
  }
}
