package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TimeTrack struct {
	gorm.Model  `gorm: `
	MethodName  string
	AppInstance string
	Time        int64
	Date        time.Time
	Result      bool
	ReasonCode  string
}

func (t TimeTrack) TableName() string {
	return "time_track"
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
