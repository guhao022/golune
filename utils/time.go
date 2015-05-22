package utils

import (
	"time"
)

func Month2Int(month time.Month) (intMonth int) {
	switch month {
	case time.January:
		intMonth = 1
	case time.February:
		intMonth = 2
	case time.March:
		intMonth = 3
	case time.April:
		intMonth = 4
	case time.May:
		intMonth = 5
	case time.June:
		intMonth = 6
	case time.July:
		intMonth = 7
	case time.August:
		intMonth = 8
	case time.September:
		intMonth = 9
	case time.October:
		intMonth = 10
	case time.November:
		intMonth = 11
	case time.December:
		intMonth = 12
	}
	return
}

func Unix2Str(t int64) string {
	tt := time.Unix(t, 0)
	timeStr := tt.Format("2006-01-02 15:04:05")
	return timeStr
}

func GetYear(t int64) int {
	tt := time.Unix(t, 0)
	return tt.Year()
}

func GetMonth(t int64) int {
	tt := time.Unix(t, 0)
	return Month2Int(tt.Month())
}

func GetDay(t int64) int {
	tt := time.Unix(t, 0)
	return tt.Day()
}
