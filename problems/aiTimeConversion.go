package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a time in
// - hour AM/PM format, convert it to military (24-hour) time.
// Note:
// - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
func timeConversion(timeStr string) string {
	timeArr := strings.Split(timeStr, ":")

	if len(timeArr) != 3 {
		return "invalid time"
	}

	isPM := strings.Contains(timeArr[2], "PM")
	isAM := !isPM

	hourStr := timeArr[0]
	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		return "invalid string on hour"
	}

	min := timeArr[1]
	second := timeArr[2][0:2]

    if isAM {
        if hour == 12 {
            hour = 0
        }
    } else if isPM {
        if hour != 12 {
            hour += 12
        }
    }

    return strings.Join([]string{fmt.Sprintf("%02d", hour), min, second}, ":")
}
