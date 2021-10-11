package testhelper

import (
	"time"
)

func MockFromDateString(layout, dateData string) *time.Time {
	tm, _ := time.Parse(layout, dateData)
	loc, _ := time.LoadLocation("Asia/Bangkok")
	tIn := (tm).In(loc)
	return &tIn
}
