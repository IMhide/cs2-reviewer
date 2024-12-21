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
	Score          int            `json:"score"`
	FinancialStats FinancialStats `json:"financialStats"`
}

type Player struct {
	SteamID string `json:"steamID"`
	GameID  int    `json:"gameID"`
	Name    string `json:"name"`
}

type Round struct {
	RoundNumber   int                   `json:"roundNumber"`
	StartTime     time.Duration         `json:"startTime"`
	EndTime       time.Duration         `json:"endTime"`
	WinningReason events.RoundEndReason `json:"winningReason"`
	WinnerState   TeamState             `json:"winnerState"`
	LoserState    TeamState             `json:"loserState"`
	KillFeed      []Kill                `json:"killFeed"`
}

type Kill struct {
	HappenedAt   time.Duration `json:"happenedAt"`
	Victim       PlayerState   `json:"victim"`
	Attacker     PlayerState   `json:"attacker"`
	Assiter      PlayerState   `json:"assister"`
	WeaponName   string        `json:"weaponName"`
	WallBang     bool          `json:"wallBang"`
	IsHeadshot   bool          `json:"isHeadshot"`
	NoScope      bool          `json:"noScope"`
	ThroughSmoke bool          `json:"throughSmoke"`
}

type PlayerState struct {
	PlayerID   int     `json:"playerID"`
	PlayerName string  `json:"playerName"`
	Flashed    bool    `json:"flashed"`
	PositionX  float64 `json:"positionX"`
	PositionY  float64 `json:"positionY"`
	PositionZ  float64 `json:"positionZ"`
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

type Match struct {
	DemoInfo DemoInfo `json:"demoInfo"`
	Teams    [2]Team  `json:"teams"`
	Rounds   []Round  `json:"rounds"`
}
