package main

import (
	"log"
	"os"

	cs "cs_review/internal/counter_strike"
)

func main() {
if len(os.Args) < 2 {
		log.Panic("Missing parameters")
		os.Exit(2)
	}

  cs.DemoReader(os.Args[1])
  //gui.GUI()
}
