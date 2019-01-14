package main

import (
	"log"
	"time"
)

func timeTrack(start time.Time, name string) int64 {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	return elapsed.Nanoseconds() / 1000000
}
