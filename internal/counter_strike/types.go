package counter_strike

import (
	"time"

	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type TeamState struct {
	ID             int            `json:"id"`
	score          int            `json:"score"`
	financialStats FinancialStats `json:"financialStats"`
}

type Player struct {
	SteamID string `json:"steamID"`
	GameID  int    `json:"gameID"`
	Name    string `json:"name"`
}

type Round struct {
	roundNumber   int                   `json:"roundNumber"`
	StartTime     time.Duration         `json:"startTime"`
	EndTime       time.Duration         `json:"endTime"`
	WinningReason events.RoundEndReason `json:"winningReason"`
	WinnerState   TeamState             `json:"winnerState"`
	LoserState    TeamState             `json:"loserState"`
	KillFeed      []Kill                `json:"killFeed"`
}

type Kill struct {
	HappenedAt time.Duration `json:"happenedAt"`
	VictimID   int           `json:"victimID"`
	AttackerID int           `json:"attackerID"`
	WeaponName string        `json:"weaponName"`
	IsHeadshot bool          `json:"isHeadshot"`
	WallBang   bool          `json:"wallBang"`
}

type DemoInfo struct {
	MapName         string        `json:"mapName"`
	ServerName      string        `json:"serverName"`
	DurationInSec   time.Duration `json:"durationInSec"`
	DurationInTick  int           `json:"durationInTick"`
	DurationInFrame int           `json:"durationInFrame"`
}

type FinancialStats struct {
	teamID        int `json:"teamID"`
	RoundStartEco int `json:"roundStartEco"`
	FreezeTimeEco int `json:"freezeTimeEco"`
	EndRoundEco   int `json:"endRoundEco"`
	SpentMoney    int `json:"spentMoney"`
}
