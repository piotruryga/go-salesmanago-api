package main

import (
	"time"
)

type TimeTrack struct {
	MethodName  string
	AppInstance string
	Time        int64
	Date        time.Time
	Result      bool
	ReasonCode  string
}

func New(methodName string, appInstance string, timing int64, result bool, reasonCode string) TimeTrack {
	return TimeTrack{
		MethodName:  methodName,
		AppInstance: appInstance,
		Time:        timing,
		Date:        time.Now(),
		Result:      result,
		ReasonCode:  reasonCode,
	}

}
