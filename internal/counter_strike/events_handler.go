package counter_strike

import (
	"fmt"

	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

//
// Events Handlers
//

//Kills Event

func KillEventHandler() func (e events.Kill) {
  return func (e events.Kill) {
		var hs string
		if e.IsHeadshot {
			hs = " (HS)"
		}
		var wallBang string
		if e.PenetratedObjects > 0 {
			wallBang = " (WB)"
		}
		fmt.Printf("%s <%v%s%s> %s\n", e.Killer, e.Weapon, hs, wallBang, e.Victim)
	}
}

// Bomb Events

func BombPlantingEventHandler() func (e events.BombPlantBegin) {
  return func (e events.BombPlantBegin){
   fmt.Println(e.Player.Name, "is planting the bomb")
  }
}

func BombPlantCancelEventHandler() func (e events.BombPlantAborted) {
  return func (e events.BombPlantAborted) {
   fmt.Println(e.Player.Name, "stoped planting")
  }
}

func BombPlantedEventHandler() func (e events.BombPlanted) {
  return func (e events.BombPlanted) {
   fmt.Println("Bomb has been planted by", e.Player.Name)
  }
}

func BombDefusingEventHandler() func (e events.BombDefuseStart) {
  return func (e events.BombDefuseStart){
    fmt.Println(e.Player.Name, "is defusing the bomb - kit ?", e.HasKit)
  }
}

func BombDefuseCancelEventHandler() func (e events.BombDefuseAborted) {
  return func (e events.BombDefuseAborted) {
   fmt.Println(e.Player.Name, "stoped defusing")
  }
}

func BombDefusedEventHandler() func (e events.BombDefused) {
  return func (e events.BombDefused) {
   fmt.Println("Bomb has been defused by", e.Player.Name)
  }
}

func BombDroppedEventHandler() func (e events.BombDropped) {
  return func (e events.BombDropped) {
   fmt.Println(e.Player.Name, "dropped the Bomb")
  }
}

func BombPickedupEventHandler() func (e events.BombPickup) {
  return  func (e events.BombPickup) {
   fmt.Println(e.Player.Name, "picked up the Bomb")
  }
}

// Rounds Event

func RoundStartEventHandler() func (e events.RoundStart){
  return func (e events.RoundStart){
    fmt.Println("--- Round Start ---")
  }
}

func RoundFinishedEventHandler() func (e events.RoundEnd){
  return  func (e events.RoundEnd){
    fmt.Println("--- Round is over ---", humanRoundEndReason(e.Reason))
    
  }
}

func humanRoundEndReason(reason events.RoundEndReason) string{
  humanReason := map[events.RoundEndReason]string{
   1: "Bomb has exploded (T WIN)",
   2: "N.A",
   3: "N.A",
   4: "Terrorist Escaped",
   5: "CT Stopped T Escape",
   6: "Terrorist Stopped",
   7: "Bomb has been defused (CT WIN)",
   8: "Counter Terrorist Win",
   9: "Terrorist Win",
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
