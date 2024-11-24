package counter_strike

import (
	"time"

	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type Team struct {
	ID   int
	Name string
}

type TeamState struct {
	ID             int
	score          int
	financialStats FinancialStats
}

type Round struct {
	roundNumber   int
	StartTime     time.Duration
	EndTime       time.Duration
	WinningReason events.RoundEndReason
	WinnerState   TeamState

	LoserState TeamState
}

type DemoInfo struct {
	MapName         string
	ServerName      string
	DurationInSec   time.Duration
	DurationInTick  int
	DurationInFrame int
}

type FinancialStats struct {
	teamID        int
	RoundStartEco int
	FreezeTimeEco int
	EndRoundEco   int
	SpentMoney    int
}
