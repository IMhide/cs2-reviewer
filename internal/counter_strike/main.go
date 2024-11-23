package counter_strike

import (
	"fmt"
	"log"
	"os"
	"time"

	libDem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
)

func DemoReader(demoPath string) {
	file, err := os.Open(demoPath)
	if err != nil {
		log.Panic("failed to open demo file: ", err)
		os.Exit(2)
	}
	defer file.Close()

	demoParser := libDem.NewParser(file)
	defer demoParser.Close()

	//
	//  DATA STRUCTURES
	//

	var teams [2]Team
	var rounds []Round
	var demoInfo DemoInfo

	//
	//  VARIABLES
	//

	var currentDuration time.Duration
	var currentGameState libDem.GameState

	//
	//  EVENTS HANDLERS
	//

	demoParser.RegisterEventHandler(RoundStartEventHandler(&rounds, &currentDuration))
	demoParser.RegisterEventHandler(RoundFinishedEventHandler(&rounds, &currentDuration))
	demoParser.RegisterEventHandler(MatchStartEventHandler(&teams, &currentGameState))

	//
	// MAIN LOOP
	//

	nextFrame, err := demoParser.ParseNextFrame()
	for nextFrame {
		currentDuration = demoParser.CurrentTime()

		nextFrame, err = demoParser.ParseNextFrame()
		if err != nil {
			log.Panic("failed to parse demo: ", err)
		}
	}

	demoHeader := demoParser.Header()
	demoInfo.MapName = demoHeader.MapName
	demoInfo.ServerName = demoHeader.ServerName
	demoInfo.DurationInSec = demoParser.CurrentTime()
	demoInfo.DurationInTick = demoParser.GameState().IngameTick()
	demoInfo.DurationInFrame = demoParser.CurrentFrame()

	fmt.Println(demoInfo)
	fmt.Println(teams)
	fmt.Println(rounds)
}
