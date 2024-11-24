package counter_strike

import (
	"time"

	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type Team struct {
	ID      int
	Name    string
	Players []Player
}

type TeamState struct {
	ID             int
	score          int
	financialStats FinancialStats
}

type Player struct {
	SteamID string
	GameID  int
	Name    string
}

type Round struct {
	roundNumber   int
	StartTime     time.Duration
	EndTime       time.Duration
	WinningReason events.RoundEndReason
	WinnerState   TeamState
	LoserState    TeamState
	KillFeed      []Kill
}

type Kill struct {
	VictimID   int
	AttackerID int
	WeaponName string
	IsHeadshot bool
	WallBang   bool
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
