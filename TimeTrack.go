package main

import "time"

type TimeTrack struct {
	MethodName  string
	AppInstance string
	Time        int64
	Date        time.Time
}
