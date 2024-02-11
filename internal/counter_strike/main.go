package counter_strike

import (
	"log"
	"os"

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

	demoParser.RegisterEventHandler(BombPlantingEventHandler())
	demoParser.RegisterEventHandler(BombPlantCancelEventHandler())
	demoParser.RegisterEventHandler(BombPlantedEventHandler())
	demoParser.RegisterEventHandler(BombDefusingEventHandler())
	demoParser.RegisterEventHandler(BombDefuseCancelEventHandler())
	demoParser.RegisterEventHandler(BombDefusedEventHandler())
	demoParser.RegisterEventHandler(BombPickedupEventHandler())
	demoParser.RegisterEventHandler(BombDroppedEventHandler())

	demoParser.RegisterEventHandler(RoundStartEventHandler())
	demoParser.RegisterEventHandler(RoundFinishedEventHandler())

	demoParser.RegisterEventHandler(KillEventHandler())

	err = demoParser.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}
}
