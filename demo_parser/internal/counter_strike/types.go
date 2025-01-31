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
	ID             int            `json:"teamId"`
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
	WallBang     bool          `json:"isWallBang"`
	IsHeadshot   bool          `json:"isHeadshot"`
	NoScope      bool          `json:"isNoScope"`
	ThroughSmoke bool          `json:"isThroughSmoke"`
}

type PlayerState struct {
	PlayerID   int     `json:"playerGameId"`
	PlayerName string  `json:"playerName"`
	Flashed    bool    `json:"isFlashed"`
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
	TeamID        int `json:"teamId"`
	RoundStartEco int `json:"roundStartMoney"`
	FreezeTimeEco int `json:"freezeTimeMoney"`
	EndRoundEco   int `json:"endRoundMoney"`
	SpentMoney    int `json:"spentMoney"`
}

type Match struct {
	DemoInfo DemoInfo `json:"demoInfo"`
	Teams    [2]Team  `json:"teams"`
	Rounds   []Round  `json:"rounds"`
}
