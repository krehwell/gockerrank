package problems

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a time in
// - hour AM/PM format, convert it to military (24-hour) time.
// Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
func timeConversion(timeStr string) string {
	timeArr := strings.Split(timeStr, ":")

	if len(timeArr) != 3 {
		return "invalid time"
	}

	isDay := strings.Contains(timeArr[2], "PM")

	hour := timeArr[0]
	min := timeArr[1]
	second := timeArr[2][0:2]

	if !isDay {
		return strings.Join([]string{hour, min, second}, ":")
	} else {
		parsedHour, err := strconv.Atoi(hour)

		if err != nil {
			return "invalid string"
		}

		amTo24Format := parsedHour + 12

		return strings.Join([]string{fmt.Sprint(amTo24Format), min, second}, ":")
	}
}
