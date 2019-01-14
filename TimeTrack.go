package main

import "time"

type TimeTrack struct {
	MethodName  string
	AppInstance string
	Time        int64
	Date        time.Time
}

func New(methodName string, appInstance string, timing int64) TimeTrack {
	return TimeTrack{
		MethodName:  methodName,
		AppInstance: appInstance,
		Time:        timing,
		Date:        time.Now()}

}
