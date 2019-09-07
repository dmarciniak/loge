package datextractor

import (
	"math"
	"regexp"
	"strconv"
	"time"
)

var (
	rgx, _ = regexp.Compile(`[0-9]{4}[-][0-9]{2}[-][0-9]{2}[T\s][0-9]{2}:[0-9]{2}:[0-9]{2}([.,][0-9]{0,3})?`)
)

// Extract method extracted date from string
// return: raw date string, date in Time and bool information about successful
func Extract(str string) (string, time.Time, bool) {
	timeStr := rgx.FindString(str)
	if timeStr == "" {
		return "", time.Time{}, false
	}

	year, _ := strconv.Atoi(timeStr[0:4])
	month, _ := strconv.Atoi(timeStr[5:7])
	day, _ := strconv.Atoi(timeStr[8:10])

	hour, _ := strconv.Atoi(timeStr[11:13])
	minute, _ := strconv.Atoi(timeStr[14:16])
	second, _ := strconv.Atoi(timeStr[17:19])

	milis := 0
	if len(timeStr) >= 21 {
		milisStr := timeStr[20:]
		milisLen := len(milisStr)
		milis, _ = strconv.Atoi(timeStr[20:])
		milis *= int(math.Pow10(9 - milisLen))
	}

	return timeStr, time.Date(year, time.Month(month), day, hour, minute, second, milis, time.UTC), true
}
