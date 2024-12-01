package counter_strike

import (
	"time"

	libDem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	libCommon "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

//
// Events Handlers
//

// Rounds Event

func RoundStartEventHandler(rounds *[]Round, currentDuration *time.Duration) func(e events.RoundStart) {
	return func(e events.RoundStart) {
		*rounds = append(*rounds, Round{StartTime: *currentDuration, roundNumber: len(*rounds)})
	}
}

func RoundFinishedEventHandler(rounds *[]Round, currentDuration *time.Duration) func(e events.RoundEnd) {
	return func(e events.RoundEnd) {
		roundLen := len(*rounds)
		(*rounds)[(roundLen - 1)].EndTime = *currentDuration
		(*rounds)[(roundLen - 1)].WinningReason = e.Reason
		(*rounds)[(roundLen - 1)].WinnerState = TeamState{
			e.WinnerState.ID(),
			e.WinnerState.Score(),
			FinancialStats{
				e.WinnerState.ID(),
				e.WinnerState.RoundStartEquipmentValue(),
				e.WinnerState.FreezeTimeEndEquipmentValue(),
				e.WinnerState.CurrentEquipmentValue(),
				e.WinnerState.MoneySpentThisRound(),
			},
		}
		(*rounds)[(roundLen - 1)].LoserState = TeamState{
			e.LoserState.ID(),
			e.LoserState.Score(),
			FinancialStats{
				e.LoserState.ID(),
				e.LoserState.RoundStartEquipmentValue(),
				e.LoserState.FreezeTimeEndEquipmentValue(),
				e.LoserState.CurrentEquipmentValue(),
				e.LoserState.MoneySpentThisRound(),
			},
		}
	}
}

func MatchStartEventHandler(teams *[2]Team, currentGameState *libDem.GameState) func(e events.MatchStart) {
	return func(e events.MatchStart) {
		terrosistTeam := (*currentGameState).TeamTerrorists()
		counterTerroristTeam := (*currentGameState).TeamCounterTerrorists()
		(*teams)[0] = Team{ID: counterTerroristTeam.ID(), Name: counterTerroristTeam.ClanName()}
		(*teams)[1] = Team{ID: terrosistTeam.ID(), Name: terrosistTeam.ClanName()}

		for _, player := range counterTerroristTeam.Members() {
			(*teams)[0].Players = append((*teams)[0].Players, Player{GameID: (*player).UserID, Name: (*player).Name})
		}

		for _, player := range terrosistTeam.Members() {
			(*teams)[1].Players = append((*teams)[1].Players, Player{GameID: (*player).UserID, Name: (*player).Name})
		}
	}
}

func KillEventHandler(rounds *[]Round, currentDuration *time.Duration) func(e events.Kill) {
	return func(e events.Kill) {
		currentRound := len(*rounds)
		if currentRound == 0 {
			return
		}

		(*rounds)[currentRound-1].KillFeed = append((*rounds)[currentRound-1].KillFeed, Kill{
			*currentDuration,
			getPlayerId(e.Victim),
			getPlayerId(e.Killer),
			e.Weapon.String(),
			e.IsHeadshot,
			e.PenetratedObjects > 1,
		})
	}
}

func getPlayerId(player *libCommon.Player) int {
	if player != nil {
		return player.UserID
	}
	return 0
}

func humanRoundEndReason(reason events.RoundEndReason) string {
	humanReason := map[events.RoundEndReason]string{
		1:  "Bomb has exploded (T WIN)",
		2:  "N.A",
		3:  "N.A",
		4:  "Terrorist Escaped",
		5:  "CT Stopped T Escape",
		6:  "Terrorist Stopped",
		7:  "Bomb has been defused (CT WIN)",
		8:  "Counter Terrorist Win",
		9:  "Terrorist Win",
		10: "Draw",
		11: "N.A",
		12: "Target Saved",
		13: "N.A",
		14: "Terrorist not Escaped",
		15: "N.A",
		16: "Game Restarting",
		17: "Terrorist Surrender",
		18: "Counter Terrorist Surrender",
	}

	return humanReason[reason]
}
