package counter_strike

import "time"

type Team struct {
	ID   int
	Name string
}

type Round struct {
	roundNumber int
	StartTime   time.Duration
	EndTime     time.Duration
}

type DemoInfo struct {
	MapName         string
	ServerName      string
	DurationInSec   time.Duration
	DurationInTick  int
	DurationInFrame int
}
