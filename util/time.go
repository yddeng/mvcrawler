package util

import "time"

var WeekDay = map[string]int{
	"Sunday":    0,
	"Monday":    1,
	"Tuesday":   2,
	"Wednesday": 3,
	"Thursday":  4,
	"Friday":    5,
	"Saturday":  6,
}

func Now() time.Time {
	return time.Now()
}

func NowUnix() int64 {
	return Now().Unix()
}

func NowUnixNano() int64 {
	return Now().UnixNano()
}

func GetWeekDay() int {
	return int(time.Now().Weekday())
}
